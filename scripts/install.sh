#!/usr/bin/env bash

if [ $EUID != 0 ]; then
    echo "please run as root"
    exit
fi

default_dir="/usr/share/void"
install_dir=${1:-default_dir}
commit=$(git log -1 | head -n 1 | sed -e "s/commit//")

if [ ! -d $install_dir ]; then
    echo "moving files..."
    mkdir /usr/share/void
    cp -r ./* /usr/share/void
    cd /usr/share/void
    rm -rf examples .gitignore .git
    echo "building..."
    go build -ldflags "-X 'utils.Commit=$commit'"
    echo "installing..."
    mv ./void /usr/local/bin
    echo "installation successful"
else
    echo "directory already exists"
    exit
fi
