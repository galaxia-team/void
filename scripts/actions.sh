#!/usr/bin/env bash

commit=$(git log -1 | head -n 1 | sed -e "s/commit//")

echo -e "building...\n"
go build -ldflags "-X 'utils.Commit=$commit'" -v
