#!/usr/bin/env bash

SCRIPTS=$(dirname $0)
SERVER_IP="updater.retrowiki.es"
SV_HOME="/home/retroserver"
SV_SERVER_NAME="retroserver"

$SCRIPTS/build.sh
if [[ $? -ne 0 ]]
then
	echo "Building error!"
	exit 1
fi

rsync \
	-avz --delete --delete-excluded --exclude=".DS_Store" \
	$SCRIPTS/../files root@$SERVER_IP:$SV_HOME/.

SERVICE=$(cat $SCRIPTS/service/retroserver.service)
FILE=$(find $SCRIPTS/../bin -name retroserver_\*-linux_amd64)
BASENAME=$(basename $FILE)
ssh root@$SERVER_IP <<-CMD
	service retroserver stop
	rm -f $SV_HOME/retroserver_v*
CMD

scp $FILE root@$SERVER_IP:$SV_HOME/.
ssh root@$SERVER_IP <<-CMD
	chmod 700 $SV_HOME/$BASENAME
	chown retroserver:retroserver $SV_HOME/$BASENAME
	chown -R retroserver:retroserver $SV_HOME/files
	ln -sf $SV_HOME/$BASENAME $SV_HOME/$SV_SERVER_NAME
	chown -h retroserver:retroserver $SV_HOME/$SV_SERVER_NAME
	echo "${SERVICE}" > /etc/systemd/system/retroserver.service
	systemctl daemon-reload
	service retroserver start
CMD

exit 0
