#!/usr/bin/bash

# 项目名称
SERVICE_NAME=$1

ps -ef | grep ${SERVICE_NAME}-http | awk '{print $2}' | xargs kill -9
ps -ef | grep ${SERVICE_NAME}-rpc | awk '{print $2}' | xargs kill -9
ps -ef | grep ${SERVICE_NAME}-back | awk '{print $2}' | xargs kill -9
ps -ef | grep ${SERVICE_NAME}-crontab | awk '{print $2}' | xargs kill -9

chmod +x /home/${SERVICE_NAME}/linux/rpc/${SERVICE_NAME}-rpc
chmod +x /home/${SERVICE_NAME}/linux/http/${SERVICE_NAME}-http
chmod +x /home/${SERVICE_NAME}/linux/back/${SERVICE_NAME}-back
chmod +x /home/${SERVICE_NAME}/linux/crontab/${SERVICE_NAME}-crontab

/home/${SERVICE_NAME}/linux/rpc/${SERVICE_NAME}-rpc -c /etc/${SERVICE_NAME} -log /var/log/${SERVICE_NAME}/rpc.log >/var/log/${SERVICE_NAME}-rpc.log 2>&1 &
/home/${SERVICE_NAME}/linux/http/${SERVICE_NAME}-http -c /etc/${SERVICE_NAME} -log /var/log/${SERVICE_NAME}/http.log >/var/log/${SERVICE_NAME}-http.log 2>&1 &
/home/${SERVICE_NAME}/linux/back/${SERVICE_NAME}-back -c /etc/${SERVICE_NAME} -log /var/log/${SERVICE_NAME}/back.log >/var/log/${SERVICE_NAME}-back.log 2>&1 &
/home/${SERVICE_NAME}/linux/crontab/${SERVICE_NAME}-crontab -c /etc/${SERVICE_NAME} -log /var/log/${SERVICE_NAME}/crontab.log >/var/log/${SERVICE_NAME}-crontab.log 2>&1 &
