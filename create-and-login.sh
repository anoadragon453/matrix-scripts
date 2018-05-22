username="testing-$(pwgen 4 1)"
password="testing123"

reg_res="$(POST '/_matrix/client/r0/register' '{"username": "'$username'", "password": "'$password'"}' 2>&1)"
session="$(echo $reg_res | jq .session)"
reg_res2="$(POST '/_matrix/client/r0/register' '{"username": "'$username'", "password": "'$password'", "auth": {"type": "m.login.dummy", "session": '$session'}}' 2>&1)"
echo $reg_res2

# -r - Don't print with surrounding quotes
user_id=$(echo "$reg_res2" | jq -r .user_id)
access_token=$(echo "$reg_res2" | jq -r .access_token)
device_id=$(echo "$reg_res2" | jq -r .device_id)

export USER_ID=$user_id
export ACCESS_TOKEN=$access_token
export DEVICE_ID=$device_id

echo ""
echo "\$USERNAME: $user_id"
echo "\$ACCESS_TOKEN: $access_token"
echo "\$DEVICE_ID: $device_id"
resty "http://127.0.0.1:8008" -H "Authorization: Bearer $ACCESS_TOKEN"
