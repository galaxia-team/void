#!/usr/bin/env bash

commit=$(git log -1 | head -n 1 | sed -e "s/commit//")

echo -e "building void...\n"
cd void
go build -ldflags "-X 'utils.Commit=$commit'" -v

echo -e "\nbuilding vpkg...\n"
cd ../vpkg
go build -ldflags "-X 'utils.Commit=$commit'" -v