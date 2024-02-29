#!/usr/bin/env bash

pushd net

git pull

popd

cp -r net/quic/* .

git grep golang.org/x/net\
/quic | cut -d : -f 1 | grep -v README.md | xargs -L1 sed -i -e s,golang.org/x/net\
/quic,github.com/webtransport/quic,g

go mod tidy
