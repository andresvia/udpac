#!/bin/bash
set -eu
# vars
CURRENT_TAG=$(git tag | tail -n1)
VERSION_MAJOR_MINOR="${TRAVIS_TAG:-$CURRENT_TAG}"
VERSION_RELEASE="${TRAVIS_BUILD_NUMBER:-$(date +%s)}"
VERSION="${VERSION_MAJOR_MINOR}.${VERSION_RELEASE}"
# clean
rm -rfv udpac
rm -rfv packages/*
# build
go build
# files
mkdir -pv package/deb/usr/sbin
mkdir -pv package/rpm/usr/sbin
cp -vf udpac package/deb/usr/sbin
cp -vf udpac package/rpm/usr/sbin
mv udpac packages/udpac
# package
mkdir -vp packages
fpm --package packages/ -C package/deb -s dir -t deb -n udpac -v "${VERSION}" --after-install package/deb-after-install .
fpm --package packages/ -C package/rpm -s dir -t rpm -n udpac -v "${VERSION}" --after-install package/rpm-after-install .
