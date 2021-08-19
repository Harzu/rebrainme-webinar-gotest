#!/bin/bash

FAIL_TEST=false
go test ./... -cover -count=1 || FAIL_TEST=true;

if [ "$FAIL_TEST" = true ]; then
  exit 1
fi
