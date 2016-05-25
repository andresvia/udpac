#!/bin/bash
set -eu
# vars
VERSION_MAJOR_MINOR="${TRAVIS_TAG:-2.0}"
VERSION_RELEASE="${TRAVIS_BUILD_NUMBER:-$(date +%s)}"
VERSION="${VERSION_MAJOR_MINOR}.${VERSION_RELEASE}"
# clean
rm -rfv udpac*
# build
go build
# files
mkdir -pv package/deb/usr/sbin
mkdir -pv package/rpm/usr/sbin
cp -vf udpac package/deb/usr/sbin
cp -vf udpac package/rpm/usr/sbin
# package
fpm -C package/deb -s dir -t deb -n udpac -v "${VERSION}" --after-install package/deb-after-install .
fpm -C package/rpm -s dir -t rpm -n udpac -v "${VERSION}" --after-install package/rpm-after-install .
