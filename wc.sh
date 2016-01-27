#!/bin/bash 
find . -type f | grep "\.go$" | grep -v magknot | xargs wc -l | grep total
find magknot -type f | grep "\.go$" | xargs wc -l |grep total
find . -type f | grep "\.go$" | xargs wc -l | grep total
