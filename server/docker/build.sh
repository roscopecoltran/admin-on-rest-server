#!/bin/sh
set -x
set -e

# Set temp environment vars
export APP_VCS_URI=${APP_VCS_URI:-"app"}
export GOPATH=/tmp/go
export PATH=${PATH}:${GOPATH}/bin
export BUILDPATH=${GOPATH}/src/${APP_VCS_URI}
export PKG_CONFIG_PATH="/usr/lib/pkgconfig/:/usr/local/lib/pkgconfig/"

# Install build deps
apk --no-cache --no-progress --virtual build-deps add ${APK_BUILD}

# Install go dependencies
go get github.com/Masterminds/glide
go get github.com/mitchellh/gox

# Init go environment to build your executable(s)
mkdir -p $(dirname ${BUILDPATH})
ln -s /app ${BUILDPATH}
cd ${BUILDPATH}

if [ ! -f glide.yaml ]; then
	yes no | glide create
fi
glide install --strip-vendor

cd ./cmd/${APP_EXECUTABLE}
go build 
go install

# Cleanup GOPATH
rm -r ${GOPATH}

# Remove build deps
apk --no-cache --no-progress del build-deps