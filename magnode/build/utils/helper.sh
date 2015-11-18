#!/bin/sh


## Macro
RELEASE=0
DEBUG=1
CLEAN=2

function usage()
{
	echo "Usage:sh build_[android|ios].sh [release|debug|clean]"
	echo "Description:"
	echo "    release/debug: release or debug library building"
	echo "    clean : clean project"
	exit 121
}

function parse_args()
{
    if [ $# != 1 ] ; then
	usage 
    fi
    if [ $1 == "release" ]; then
	return $RELEASE
    elif [ $1 == "debug" ]; then
	return $DEBUG
    elif [ $1 == "clean" ]; then
	return $CLEAN
    else
	usage
    fi
}

STLS=("c11" "c98")	
function valied_stl()
{

	for i in ${STLS[@]} 
	do
		if [ $1 == ${i}  ] ; then
			return 1
		fi
	done
	return 0
}

function succ_exit()
{
    time_str=`date`
    echo "[SUCCESS("$time_str")] Done Build !"
    exit 0
}


function fail_exit()
{
    time_str=`date`
    echo "[FAILED("$time_str")] Build  Failed!"
    exit 120
}

function cancel_exit()
{
    time_str=`date`
    echo "[SUCCESS("$time_str")] Clean  Success!"
    exit 0
}
