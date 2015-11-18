#!/bin/sh

LSLOGSTACK () {
local i=0
local FRAMES=${#BASH_LINENO[@]}
# FRAMES-2 skips main, the last one in arrays
for ((i=FRAMES-2; i>=0; i--)); do
echo '  File' \"${BASH_SOURCE[i+1]}\", line ${BASH_LINENO[i]}, in ${FUNCNAME[i+1]}
# Grab the source code of the line
#sed -n "${BASH_LINENO[i]}{s/^/    /;p}" "${BASH_SOURCE[i+1]}"
done
}

function runcmd()
{
#local s=$*
$*
local ret=$?
if [ $ret -eq 0 ]
then
local a="done"
#echo $ret
#echo found
else
echo $ret
echo `pwd`"/"$0":"$LINENO":Error: "$*
LSLOGSTACK
exit $ret
fi
}
