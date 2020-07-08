#!/usr/bin/env bash


CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o ./build/imageMerge_win32/editer_pro.exe main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/imageMerge_win64/editer_pro.exe main.go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/imageMerge_mac/editer_pro main.go

rm -f ./imageMerge_win32.zip
rm -f ./imageMerge_win64.zip
rm -f ./imageMerge_mac.zip

zip -q -r imageMerge_win32.zip ./build/imageMerge_win32
zip -q -r imageMerge_win64.zip ./build/imageMerge_win64
zip -q -r imageMerge_mac.zip ./build/imageMerge_mac
