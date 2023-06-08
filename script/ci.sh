#!/bin/bash
set -eu

repo_dir="$(git rev-parse --show-toplevel)"

chmod o+w "${repo_dir}/test-data/logs" # for mysql container
docker compose up -d db

cat <<EOF > .env
MYSQL_HOST=localhost
MYSQL_DATABASE=test
MYSQL_USER=root
MYSQL_PASSWORD=root
EOF
bash "${repo_dir}/script/wait-db.sh"
go test -cover "${repo_dir}/internal/..."

docker compose down
