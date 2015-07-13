#!/usr/bin/env bash

rename -s cdnv_ mn_ cdnv_*

sed -i .bak 's/cdnv_/mn_/g;s/CDNV_/MN_/g;s/ apollo / cz /;s/ apollo\./cz\./;s/cdnvister/magnode/' mn_*.c
sed -i .bak 's/cdnv_/mn_/g;s/CDNV_/MN_/g;s/ apollo / cz /;s/ apollo\./cz\./;s/cdnvister/magnode/' mn_*.h
rm -rf *.bak

