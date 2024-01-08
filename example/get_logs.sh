#!/bin/bash

LIMIT=$1

logfiles=$(ls ./logs | jq -R | jq -s | jq ".[:$LIMIT]")

echo "{ \"logfiles\": $logfiles }"
