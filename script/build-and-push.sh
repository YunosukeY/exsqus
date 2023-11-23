#!/bin/bash
set -eu

function split_version() {
  local version="$1"
  local vs="${version//./ }"
  major=${vs[0]}
  minor=${vs[1]}
  patch=${vs[2]}
}

split_version "$1"
tags=("latest" "${major}" "${major}.${minor}" "${major}.${minor}.${patch}")

docker login -u kimitsu -p "${DOCKERHUB_PASSWORD}"

for t in "${tags[@]}"
do
  echo "$t"
  docker build -t "kimitsu/exsqus:${t}" .
  docker push "kimitsu/exsqus:${t}"
done
