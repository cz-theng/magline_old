#!/usr/bin/env sh
make clean
make
rm -rf log

./chatroom -c config/config.json  -cpuprof chatroom.prof 
