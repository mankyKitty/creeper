#!/bin/bash

CWD=$(pwd)

cd /Users/manky/gocode/src/github.com/mankyKitty/creeper

echo "Building..."
GOOS=windows GOARCH=386 go build -o _crosscomp/win86/creeper.exe creeper.go
GOOS=linux GOARCH=386 go build -o _crosscomp/linux86/creeper creeper.go
GOOS=linux GOARCH=amd64 go build -o _crosscomp/linux64/creeper creeper.go
go build -o _crosscomp/osx64/creeper creeper.go
echo "Finished building"

echo "Packaging..."
ls -l1 _crosscomp/ | while read a; do tar -czf creeper-$a-v0-1-alpha.tar.gz _crosscomp/$a;done
echo "Finished Packging"

cd $CWD