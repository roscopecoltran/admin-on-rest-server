#!/bin/sh
set -x
set -e

case "$1" in
  'run-build')
	export APP_PORT=${APP_PORT:-"3000"}  
	yarn global add serve
	npm run-script build
  	exec serve -s build -p ${APP_PORT} $@
	;;
  'bash')
	apk add --no-cache --no-progress ${APK_DEV}
  	exec /bin/bash $@
  	;;
  *)
  	exec npm run
	;;
esac