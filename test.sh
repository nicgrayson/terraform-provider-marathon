#!/bin/bash

if [ -z "$TF_VAR_marathon_url" ]; then
  TF_ACC=yes MARATHON_URL=$TF_VAR_marathon_url go test ./marathon -v
else
  echo "Please set the TF_VAR_marathon_url environment variable"
fi
