#!/bin/bash

FAIL_TEST=false

for PACKAGE in $(go list ./internal/...); do
    go test $PACKAGE -cover -count=1 || FAIL_TEST=true;
done

if [ "$FAIL_TEST" = true ]; then
  exit 1
fi
