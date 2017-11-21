#!/bin/sh
set -x
set -e

function link_go_source() {
    export BUILDPATH=${GOPATH}/src/${APP_VCS_URI}    
    mkdir -p $(dirname ${BUILDPATH})
    ln -s /app ${BUILDPATH}
    cd ${BUILDPATH}
}

case "$1" in
  'dev')
    export APK_BUILD=${APK_BUILD:-"git libssh2 go gcc musl-dev make cmake openssl-dev libssh2-dev"}
    export APK_DEV=${APK_DEV:-"nano bash jq tree"}
    export PKG_CONFIG_PATH="/usr/lib/pkgconfig/:/usr/local/lib/pkgconfig/"
    apk add --no-cache --no-progress ${APK_BUILD} ${APK_DEV}
    link_go_source()
    exec /bin/bash # $@
    ;;
  'bash')
	  apk add --no-cache --no-progress bash
  	exec /bin/bash # $@
  	;;
  *)
    exec /app/dist/${APP_EXECUTABLE} $@
	;;
esac