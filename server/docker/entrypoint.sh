#!/bin/sh
set -x
set -e

function symlink_go_source() {
    export BUILDPATH=${GOPATH}/src/${APP_VCS_URI}    
    mkdir -p $(dirname ${BUILDPATH})
    ln -s /app ${BUILDPATH}
}

case "$1" in
  'dev')
    export GOPATH=/usr/share/go
    export PATH=${PATH}:${GOPATH}/bin
    export APK_BUILD=${APK_BUILD:-"git libssh2 go gcc musl-dev make cmake openssl-dev libssh2-dev"}
    export APK_DEV=${APK_DEV:-"nano bash jq tree"}
    export PKG_CONFIG_PATH="/usr/lib/pkgconfig/:/usr/local/lib/pkgconfig/"
    apk add --no-cache --no-progress ${APK_BUILD} ${APK_DEV}
    symlink_go_source
    go get github.com/Masterminds/glide
    go get github.com/mitchellh/gox
    exec /bin/bash
    ;;
  'bash')
    export APK_DEV=${APK_DEV:-"nano bash jq tree"}
	  apk add --no-cache --no-progress ${APK_DEV}
  	exec /bin/bash
  	;;
  *)
    exec /app/dist/${APP_EXECUTABLE} $@
	;;
esac