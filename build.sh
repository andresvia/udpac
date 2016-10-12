#!/bin/bash
set -eu
# use gox
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
gox -os 'Linux Windows' -arch 'amd64 386'
package_artifact="udpac_linux_amd64"
artifacts="udpac_linux_386 udpac_linux_amd64 udpac_windows_386.exe udpac_windows_amd64.exe"
# files
mkdir -pv package/os/usr/sbin
cp -vf $package_artifact package/os/usr/sbin/udpac
# packages and binary
mkdir -vp packages
mv $artifacts packages/
fpm --package=packages/ -C package/os -s dir -t deb --name=udpac --version="${VERSION}" --after-install package/after-install --deb-no-default-config-files .
fpm --package=packages/ -C package/os -s dir -t rpm --name=udpac --version="${VERSION}" --after-install package/after-install --rpm-os linux .
