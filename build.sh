#!/bin/bash
set -eu
# detect os
if [ "$(uname -sm)" = "Linux x86_64" ]
then
  go_build="go build"
  artifact="udpac"
else
  go_build="gox -os Linux -arch amd64"
  artifact="udpac_linux_amd64"
fi
# vars
CURRENT_TAG=$(git tag | tail -n1)
VERSION_MAJOR_MINOR="${TRAVIS_TAG:-$CURRENT_TAG}"
VERSION_RELEASE="${TRAVIS_BUILD_NUMBER:-$(date +%s)}"
VERSION="${VERSION_MAJOR_MINOR}.${VERSION_RELEASE}"
# clean
rm -rfv udpac udpac_linux_amd64
rm -rfv packages/*
# get pkgs
go get -t -v ./...
# build
$go_build
# files
mkdir -pv package/os/usr/sbin
cp -vf $artifact package/os/usr/sbin/udpac
mv $artifact packages/udpac_linux_amd64
# package
mkdir -vp packages
fpm --package=packages/ -C package/os -s dir -t deb --name=udpac --version="${VERSION}" --after-install package/after-install --deb-no-default-config-files .
fpm --package=packages/ -C package/os -s dir -t rpm --name=udpac --version="${VERSION}" --after-install package/after-install --rpm-os linux .
