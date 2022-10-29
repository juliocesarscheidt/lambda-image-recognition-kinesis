#!/bin/bash

export SRC_PATH="${PWD}/../"

########################### sonarqube ###########################
docker container run --rm -d \
  --name sonarqube \
  --network bridge \
  -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true \
  -p 9000:9000 \
  sonarqube:lts 2> /dev/null

export SONARQUBE_URL='127.0.0.1:9000'
until [ "$(curl --silent -X GET "http://${SONARQUBE_URL}/api/system/status" 2> /dev/null | jq -r '.status')" == "UP" ]; do
  echo "Waiting for sonarqube to be up, sleeping 10 secs..."
  sleep 10
done

# change default password
# default user: admin, default password: admin
if [ -f ".sonarqube_password" ]; then
  export RANDOM_PASS=$(cat ".sonarqube_password")
else
  export RANDOM_PASS=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)
  echo "${RANDOM_PASS}" > ".sonarqube_password"
  curl --silent -X POST \
    -u 'admin:admin' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode "login=admin" \
    --data-urlencode "previousPassword=admin" \
    --data-urlencode "password=${RANDOM_PASS}" \
    --url "http://${SONARQUBE_URL}/api/users/change_password"
fi

# create a token
if [ -f ".sonarqube_token" ]; then
  export TOKEN=$(cat ".sonarqube_token")
else
  export BASIC_TOKEN="$(echo -n 'admin:'${RANDOM_PASS}'' | base64 -w0)"
  export TOKEN=$(curl --silent -X POST \
    -H "Authorization: Basic ${BASIC_TOKEN}" \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode "name=api-token" \
    --data-urlencode "login=admin" \
    --url "http://${SONARQUBE_URL}/api/user_tokens/generate" | jq -r '.token')
  echo "${TOKEN}" > ".sonarqube_token"
fi

########################### scan ###########################
function scan_project() {
  local PROJECT_KEY="$1"
  cd "${SRC_PATH}/${PROJECT_KEY}"
  go test -cover -coverpkg="github.com/juliocesarscheidt/${PROJECT_KEY}/application/usecase" -coverprofile cover.out tests/**/*_test.go
  docker container run --rm \
    --name sonarscanner \
    --network host \
    -e SONAR_HOST_URL="http://${SONARQUBE_URL}" \
    -e SONAR_SCANNER_OPTS="-Dsonar.projectKey=${PROJECT_KEY}" \
    -e SONAR_LOGIN="${TOKEN}" \
    -v "${PWD}:/usr/src" \
    -v sonar-cache:/opt/sonar-scanner/.sonar/cache \
    -w /usr/src \
    sonarsource/sonar-scanner-cli:4
}

########################### consumer ###########################
scan_project 'lambda-consumer'

########################### consumer ###########################
scan_project 'lambda-producer'
