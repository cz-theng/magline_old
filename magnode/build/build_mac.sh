#!/usr/bin/env sh
echo "Building iOS ..."
source ./utils/shutil.sh
source ./utils/helper.sh
CWD=`pwd`
parse_args $@
debug_flag=$?

mac_sdk=`xcodebuild -showsdks 2>/dev/null | grep "macosx" | awk '{print $5}'`

cd $CWD
ver=1.0.1

function copy2dist_release()
{
	echo "Copy to distribution direcoty..."
	cd $CWD
}

function copy2dist_debug()
{
	echo "Copy to distribution direcoty..."
	cd $CWD
}

function build_xcode_release()
{
	#	runcmd xcodebuild   -configuration Release -sdk ${ios_sdk} -destination 'platform=iOS,arch=\"arm64 armv7 armv7s\"'
	runcmd xcodebuild   -configuration Release -sdk ${mac_sdk} -destination 'platform=OSX'
}

function build_xcode_debug()
{
	runcmd xcodebuild   -configuration Debug  -sdk ${mac_sdk} -destination 'platform=OSX'
}

function clean_xcode()
{
	runcmd xcodebuild  clean   -configuration Debug   -sdk ${mac_sdk} -destination 'platform=OSX'
	runcmd xcodebuild  clean   -configuration Release -sdk ${mac_sdk} -destination 'platform=OSX'
}


function build_projs_debug()
{
	cd $CWD && cd ./MacOS/magnode/ && build_xcode_debug
	cd $CWD
}

function build_projs_release()
{
	cd $CWD
	cd $CWD && cd ./MacOS/magnode/ && build_xcode_release
	cd $CWD
}

function clean_projs()
{
	cd $CWD
	cd $CWD && cd ./MacOS/magnode/ && clean_xcode && rm -rf build
	cd $CWD

}

cd $CWD
cd ./MacOS/magnode/
if [ $debug_flag == $RELEASE ]; then
    echo "Build Release Library ..."
	build_projs_release
	copy2dist_release
    succ_exit
elif [ $debug_flag == $DEBUG ]; then
    echo "Build Debug Library ..."
	build_projs_debug
	copy2dist_debug
    succ_exit
elif [ $debug_flag == $CLEAN ]; then
    echo "Clean Project ..."
	clean_projs
    cancel_exit
else 
    fail_exit
fi

