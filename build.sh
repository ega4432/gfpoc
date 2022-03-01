#!/bin/sh
cd capturer
go build -o ../bin/capturer

cd ../cleaner
go build -o ../bin/cleaner

cd ../sender
go build -o ../bin/sender

cd ../
go build -o ./gfpoc
