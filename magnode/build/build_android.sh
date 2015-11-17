#!/usr/bin/env sh
source ./utils/shutil.sh
source ./utils/helper.sh

CWD=`pwd`
parse_args $@
debug_flag=$?


function copy2dist()
{
	echo "Copy to distribution direcoty..."
	cd $CWD
}


cd $CWD
ver=1.0.1 

cd $CWD
cd ./Android/jni/ 
if [ $debug_flag == $RELEASE ]; then
    echo "Build Release Library ..."
	#runcmd ndk-build V=1 NDK_DEBUG=0
    runcmd ndk-build  NDK_DEBUG=0
	copy2dist
    succ_exit
elif [ $debug_flag == $DEBUG ]; then
    echo "Build Debug Library ..."
	#runcmd ndk-build V=1 NDK_DEBUG=1
	runcmd ndk-build NDK_DEBUG=1
	copy2dist
    succ_exit
elif [ $debug_flag == $CLEAN ]; then
    echo "Clean Project ..."
    runcmd ndk-build V=1 clean
    cancel_exit
else 
    fail_exit
fi



