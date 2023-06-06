#!/usr/bin/env bash

set -eu

chmod o+w test-data/logs # for mysql container
docker compose up -d
sudo chmod o+r test-data/logs/slow.log # for app container
mysql --protocol=tcp -h localhost -P 3306 -u root -proot -e "SELECT SLEEP(2);"
logs="$(docker logs app 2>&1)"

query='"Query":"SELECT SLEEP(2);"'
if [[ "$logs" == *"$query"* ]]; then
  echo "query: ok"
else
  docker compose logs
  docker compose down
  exit 1
fi

plan='"plan":[{"Id":1,"SelectType":{"String":"SIMPLE","Valid":true},"Table":{"String":"","Valid":false},"Partitions":{"String":"","Valid":false},"AccessType":{"String":"","Valid":false},"PossibleKeys":{"String":"","Valid":false},"Key":{"String":"","Valid":false},"KeyLen":{"String":"","Valid":false},"Ref":{"String":"","Valid":false},"Rows":{"String":"","Valid":false},"Filtered":{"String":"","Valid":false},"Extra":{"String":"No tables used","Valid":true}}]'
if [[ "$logs" == *"$plan"* ]]; then
  echo "plan: ok"
else
  docker compose logs
  docker compose down
  exit 1
fi

docker compose down
