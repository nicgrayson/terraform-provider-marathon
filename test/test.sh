#!/bin/bash
set -e

go list github.com/Banno/terraform-provider-marathon \
    | xargs go list -f '{{join .Deps "\n"}}' \
    | grep -v github.com/Banno/terraform-provider-marathon \
    | sort -u \
    | xargs go get -f -u -v

go test -v .
