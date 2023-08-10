#!/bin/sh

ps -ef | grep http | awk '{print $1}' | xargs kill -9
ps -ef | grep rpc | awk '{print $1}' | xargs kill -9
ps -ef | grep back | awk '{print $1}' | xargs kill -9
ps -ef | grep crontab | awk '{print $1}' | xargs kill -9

chmod +x /deploy/rpc/rpc
chmod +x /deploy/http/http
chmod +x /deploy/back/back
chmod +x /deploy/crontab/crontab

/deploy/rpc/rpc -c /deploy -log /deploy/log/rpc.log &
/deploy/http/http -c /deploy -log /deploy/log/http.log &
/deploy/back/back -c /deploy -log /deploy/log/back.log &
/deploy/crontab/crontab -c /deploy -log /deploy/log/crontab.log &

tail -f /dev/null
