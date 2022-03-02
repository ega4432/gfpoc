#!/bin/sh
set -uex

cd ./cmd/capturer
go build -o ../../bin/capturer

cd ../cleaner
go build -o ../../bin/cleaner

cd ../notify
go build -o ../../bin/notify

# cd ../../
# go build -o ./gfpoc
