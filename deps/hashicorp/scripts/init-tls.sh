# Store root token in a file so it can be shared with other services through volume
mkdir -p $TOKEN_PATH

echo "Initializing Vault: ${VAULT_ADDR}"

# Init Vault
curl --cacert $VAULT_CACERT --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
     --data '{"secret_shares": 1, "secret_threshold": 1}' ${VAULT_ADDR}/v1/sys/init > init.json

# Retrieve root token and unseal key
VAULT_TOKEN=$(cat init.json | jq .root_token | tr -d '"')
UNSEAL_KEY=$(cat init.json | jq .keys | jq .[0])
SHA256SUM=$(cat $PLUGIN_PATH/SHA256SUM)
rm init.json

if [ "$UNSEAL_KEY" = "null" ]; then
  echo "cannot retrieve unseal token"
  exit 1
fi

echo $VAULT_TOKEN > $TOKEN_PATH/.root
echo "ROOT_TOKEN: $VAULT_TOKEN"
echo "SHA256SUM: ${SHA256SUM}"

# Unseal Vault
curl --cacert $VAULT_CACERT --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
     --data '{"key": '${UNSEAL_KEY}'}' ${VAULT_ADDR}/v1/sys/unseal

# Enable kv-v2 secret engine
curl --cacert $VAULT_CACERT --header "X-Vault-Token: ${VAULT_TOKEN}" --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
        --data '{"type": "kv-v2", "config": {"force_no_cache": true} }' \
    ${VAULT_ADDR}/v1/sys/mounts/secret


# Register Quorum plugin
curl --cacert $VAULT_CACERT --header "X-Vault-Token: ${VAULT_TOKEN}" --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  --data "{\"sha256\": \"${SHA256SUM}\", \"command\": \"quorum\" }" \
  ${VAULT_ADDR}/v1/sys/plugins/catalog/secret/quorum

# Enable quorum secret engine
curl --cacert $VAULT_CACERT --header "X-Vault-Token: ${VAULT_TOKEN}" --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  --data '{"type": "plugin", "plugin_name": "quorum", "config": {"force_no_cache": true, "passthrough_request_headers": ["X-Vault-Namespace"]} }' \
  ${VAULT_ADDR}/v1/sys/mounts/quorum

# Enable role policies
# Instructions taken from https://learn.hashicorp.com/tutorials/vault/getting-started-apis
curl --cacert $VAULT_CACERT --header "X-Vault-Token: ${VAULT_TOKEN}" --request POST \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  --data '{"type": "approle"}' \
  ${VAULT_ADDR}/v1/sys/auth/approle

curl --cacert $VAULT_CACERT --header "X-Vault-Token: $VAULT_TOKEN" \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  --request PUT \
  --data '{ "policy":"path \"quorum/*\" { capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\"] }" }' \
  ${VAULT_ADDR}/v1/sys/policies/acl/allow_secrets

curl --cacert $VAULT_CACERT --header "X-Vault-Token: $VAULT_TOKEN" \
  --request POST \
  --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  --data '{"policies": ["allow_secrets"]}' \
  ${VAULT_ADDR}/v1/auth/approle/role/key-manager

curl --cacert $VAULT_CACERT --header "X-Vault-Token: $VAULT_TOKEN" \
     --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  ${VAULT_ADDR}/v1/auth/approle/role/key-manager/role-id >role.json
ROLE_ID=$(cat role.json | jq .data.role_id | tr -d '"')
echo $ROLE_ID > $TOKEN_PATH/role

curl --cacert $VAULT_CACERT --header "X-Vault-Token: $VAULT_TOKEN" \
  --request POST \
  --cert $VAULT_CLIENT_CERT --key $VAULT_CLIENT_KEY \
  ${VAULT_ADDR}/v1/auth/approle/role/key-manager/secret-id >secret.json
SECRET_ID=$(cat secret.json | jq .data.secret_id | tr -d '"')
echo $SECRET_ID > $TOKEN_PATH/secret