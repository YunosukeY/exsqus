#!/usr/bin/env bash
set -eu

usage() {
  cat <<USAGE
  Usage:
  - kind-e2e.sh create
  - kind-e2e.sh run
  - kind-e2e.sh delete
USAGE
}

if [ "$#" != 1 ]; then
  usage
  exit 1
fi

command="$1"
repo_dir="$(git rev-parse --show-toplevel)"

if [ "$command" == "create" ]; then
  kind create cluster --config "${repo_dir}/kind/cluster.yaml"
  kubectl apply -k "${repo_dir}/kind"
elif [ "$command" == "run" ]; then
  mysql --protocol=tcp -h localhost -P 3306 -u root -proot -e "SELECT SLEEP(2);"

  logs="$(kubectl logs -n app db -c exsqus)"

  query='"Query":"SELECT SLEEP(2);"'
  if [[ "$logs" == *"$query"* ]]; then
    echo "query: ok"
  else
    exit 1
  fi

  plan='"Rows":[{"Id":1,"SelectType":{"String":"SIMPLE","Valid":true},"Table":{"String":"","Valid":false},"Partitions":{"String":"","Valid":false},"AccessType":{"String":"","Valid":false},"PossibleKeys":{"String":"","Valid":false},"Key":{"String":"","Valid":false},"KeyLen":{"String":"","Valid":false},"Ref":{"String":"","Valid":false},"Rows":{"String":"","Valid":false},"Filtered":{"String":"","Valid":false},"Extra":{"String":"No tables used","Valid":true}}]'
  if [[ "$logs" == *"$plan"* ]]; then
    echo "plan: ok"
  else
    exit 1
  fi
elif [ "$command" == "delete" ]; then
  kind delete cluster
else
  usage
  exit 1
fi
