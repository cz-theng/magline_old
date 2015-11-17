#!/usr/bin/env sh
source ./utils/shutil.sh
source ./utils/helper.sh

CWD=`pwd`
cd $CWD
runcmd sh build_android.sh release
runcmd sh build_ios.sh release
