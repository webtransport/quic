#!/usr/bin/env bash

pushd net

git pull

popd

cp -r net/internal/quic/* .

git grep golang.org/x/net\
/internal/quic | cut -d : -f 1 | xargs -L1 sed -i -e s,golang.org/x/net\
/internal/quic,github.com/webtransport/quic,g

go mod tidy
