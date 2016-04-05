#!/bin/bash

TF_LOG=TRACE TF_LOG_PATH=./test-sh-tf.log TF_ACC=yes MARATHON_URL=${MARATHON_URL:="https://marathon.dev.banno.com"} go test ./marathon -v
