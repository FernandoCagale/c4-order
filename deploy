#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

echo "\n===> Generate image...\n"

docker build --no-cache -t c4-order .

echo "\n===> Docker tag...\n"

docker tag c4-order fernandocagale/c4-order

echo "\n===> Docker push...\n"

docker push fernandocagale/c4-order