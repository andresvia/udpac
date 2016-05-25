#!/bin/bash
set -eu
fpm --verbose -C package/deb -s dir -t deb -n udpac -v "$TRAVIS_TAG" --after-install package/deb-after-install .
fpm --verbose -C package/rpm -s dir -t rpm -n udpac -v "$TRAVIS_TAG" --after-install package/rpm-after-install .
