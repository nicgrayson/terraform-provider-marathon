#!/bin/bash
TF_ACC=yes MARATHON_URL="http://marathon.dev.banno.com" go test ./marathon -v
