#!/bin/bash
set -eu
fpm -s dir -t deb -n udpac -v "$TRAVIS_TAG" -d metainit --after-install package/deb-after-install package/deb
