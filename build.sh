#!/bin/bash
name="Api"
strip=""
#flags=""
if [[ -z "$1" ]]
    then
        name="RemotePlayDetached"
    else
        name=$1
fi

if [[ ! -z "$2" ]]
    then
        if [ $2 = "strip" ]
            then
                strip='-ldflags=-s'
                #flags="-s"
                printf "stripping all binaries...\n"
        fi
fi

if [ ! -d "build" ]; then
  mkdir build
fi


printf "building binaries..."
CGO_ENABLED=1 fyne-cross linux .

rm -rf build/fyne-cross
mv fyne-cross build/
#printf "building linux amd64 binary..."
#GOARCH=amd64 go build -o $name $strip ..
#printf "done\n"

#printf "builiding windows amd64 binary..."
#GOARCH=amd64 GOOS=windows go build -o $name".exe" $strip ..
#printf "done\n"
