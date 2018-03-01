#!/usr/bin/env bash
set -e -u -x
# clean
artifacts="udpac_linux_386 udpac_linux_amd64 udpac_windows_386.exe udpac_windows_amd64.exe"
# shellcheck disable=SC2086
rm -rfv udpac $artifacts package/os/usr/sbin/udpac
rm -rfv packages
# get pkgs
go get -t -v ./...
# build
gox -os 'Linux Windows' -arch 'amd64 386'
# files
mkdir -pv package/os/usr/sbin
cp -vf udpac_linux_amd64 package/os/usr/sbin/udpac
# packages and binary
mkdir -vp packages
# shellcheck disable=SC2086
mv $artifacts packages/
fpm --debug --package=packages/ -C package/os -s dir -t deb --name=udpac --version="${TRAVIS_TAG:-SNAPSHOT}" --after-install package/after-install --deb-no-default-config-files .
fpm --debug --package=packages/ -C package/os -s dir -t rpm --name=udpac --version="${TRAVIS_TAG:-SNAPSHOT}" --after-install package/after-install --rpm-os linux .
