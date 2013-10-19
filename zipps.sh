#!/bin/bash

CWD=$(pwd)

cd /Users/manky/gocode/src/github.com/mankyKitty/creeper
ls -l1 _crosscomp/ | while read a; do tar -czf creeper-$a-v0-1-alpha.tar.gz _crosscomp/$a;done

cd $CWD