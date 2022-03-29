#!/usr/bin/env bash

if [ $EUID != 0 ]; then
    echo "please run as root"
    exit
fi

commit=$(git log -1 | head -n 1 | sed -e "s/commit//")

echo -e "installing void...\n"

if [ ! -d "/usr/share/void" ]; then
    echo "moving files..."
    mkdir /usr/share/void
    cp -r ./void/* /usr/share/void
    cd /usr/share/void
    echo "building..."
    go build -ldflags "-X 'utils.Commit=$commit' -X 'utils.RootDir=/usr/share/void/'"
    echo "installing..."
    mv ./void /usr/local/bin
    echo -e "void installed successfully.\n"
else
    echo "directory '/usr/share/void' already exists!"
    exit
fi

echo -e "installing vpkg...\n"

if [ ! -d "/usr/share/vpkg" ]; then
    echo "moving files..."
    mkdir /usr/share/vpkg
    cp -r ./vpkg/* /usr/share/vpkg
    cd /usr/share/vpkg
    echo "building..."
    go build -ldflags "-X 'utils.Commit=$commit' -X 'utils.RootDir=/usr/share/vpkg/'"
    echo "installing..."
    mv ./vpkg /usr/local/bin
    echo -e "vpkg installed successfully.\n"
else
    echo "directory '/usr/share/vpkg' already exists!"
    exit
fi

echo -e "installation successful."
