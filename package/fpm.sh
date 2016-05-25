#!/bin/bash
set -eu
fpm --verbose -C package/deb -s dir -t deb -n udpac -v "$TRAVIS_TAG" -d metainit --after-install package/deb-after-install
