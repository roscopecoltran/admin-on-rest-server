#!/bin/sh
set -x
set -e

function golang_deps() {
    go get github.com/Masterminds/glide
    go get github.com/mitchellh/gox
}

function glide_deps() {
    cd $1
    if [ ! -f glide.yaml ]; then
      yes no | glide create
    fi
    if [ -f glide.yaml ]; then
      glide install --strip-vendor
    fi
}

function symlink_go_source() {
    export BUILDPATH=${GOPATH}/src/${APP_VCS_URI}    
    mkdir -p $(dirname ${BUILDPATH})
    ln -s /app ${BUILDPATH}
}

case "$1" in
  'run')
    symlink_go_source
    golang_deps
    glide_deps ${BUILDPATH}
    cd ${BUILDPATH}
    exec go run cmd/aor-server-basic/main.go --port 7000
    ;;
  'dev')
    export APK_DEV=${APK_DEV:-"nano bash jq tree"}
    export PKG_CONFIG_PATH="/usr/lib/pkgconfig/:/usr/local/lib/pkgconfig/"
    apk add --no-cache --no-progress ${APK_DEV}
    symlink_go_source
    golang_deps
    glide_deps ${BUILDPATH}
    cd ${BUILDPATH}
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