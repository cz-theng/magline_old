#!/usr/bin/env sh

make clean
make 

./maglined -c config/config.json
