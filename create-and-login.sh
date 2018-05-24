username="testing-$(pwgen 4 1)"
password="testing123"

reg_res="$(curl -sL --data-binary \{\"username\":\ \"$username\",\ \"password\":\ \"$password\"\} -X POST "http://127.0.0.1:8008/_matrix/client/r0/register")"
session="$(echo $reg_res | jq -r .session)"
reg_res2="$(curl -sL --data-binary \{\"username\":\ \"$username\",\ \"password\":\ \"$password\"\,\ \"auth\":\ \{\"type\":\ \"m.login.dummy\",\ \"session\":\ \"$session\"\}\} -X POST "http://127.0.0.1:8008/_matrix/client/r0/register")"

# -r - Don't print with surrounding quotes
user_id=$(echo "$reg_res2" | jq -r .user_id)
access_token=$(echo "$reg_res2" | jq -r .access_token)
device_id=$(echo "$reg_res2" | jq -r .device_id)

export USER_ID=$user_id
export ACCESS_TOKEN=$access_token
export DEVICE_ID=$device_id
export PASSWORD=$password

echo ""
echo "\$USERNAME: $user_id"
echo "\$ACCESS_TOKEN: $access_token"
echo "\$DEVICE_ID: $device_id"
echo "\$PASSWORD: $password"
