#!/usr/bin/env bash

pushd net

git pull

popd

cp -r net/internal/quic/* .

go mod tidy
