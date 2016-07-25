#!/usr/bin/env sh
make clean
make
rm -rf log

./maglined -c config/config.json  -cpuprof echosvr.prof 
