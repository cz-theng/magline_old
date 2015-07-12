#!/usr/bin/env bash

rename -s apollo_cdn_ mn_ apollo_cdn_*
rename -s .cpp .c *.cpp

sed -i .bak 's/apollo_cdn_/mn_/;s/APOLLO_CDN_/MN_/;s/ apollo / cz /;s/ apollo\./cz\./;s/cdnvister/magnode/' mn_*.c
sed -i .bak 's/apollo_cdn_/mn_/;s/APOLLO_CDN_/MN_/;s/ apollo / cz /;s/ apollo\./cz\./;s/cdnvister/magnode/' mn_*.h
rm -rf *.bak

