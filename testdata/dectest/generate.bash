#!/usr/bin/env bash

set -xeuo pipefail

curl -sS http://speleotrove.com/decimal/dectest.zip >dectest.zip
unzip -o dectest.zip
rm dectest.zip
