/usr/bin/bash 
find . -type f | grep "\.go" | grep -v app | xargs wc -l
find ../magknot -type f | grep "\.go" | xargs wc -l
