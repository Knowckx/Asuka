#!/usr/bin/env bash

set -e

mode=$1
if [ $mode != "uat" ] && [ $mode != "prod" ] 
then
    echo "check arg"
    exit
fi


KubeCfg=uat
REMOTE_HOST=uat-02.ckisix0kw0iw.us-west-1.rds.amazonaws.com
NameSpace=aeolia-uat
REMOTE_PORT=5432
LOCAL_PORT=5433
TEMP_POD_NAME=zc-port-jump



if [ $val1 == "prod" ]
then
    echo "use prod"
    REMOTE_HOST=cd-daas-shared-01.ckisix0kw0iw.us-west-1.rds.amazonaws.com
    NameSpace=smec
    KubeCfg=eureka--cd
elif [ $val1 == "uat" ]
then
    echo "use uat"
fi

kubecm switch $KubeCfg


function cleanup {
  echo "cleanup..." 
  kubectl delete -n $NameSpace pod/$TEMP_POD_NAME --grace-period 1 --wait=false
  kubecm switch uat
}
trap cleanup EXIT

kubectl run -n $NameSpace --image marcnuri/port-forward $TEMP_POD_NAME \
    --env REMOTE_HOST=$REMOTE_HOST \
    --env REMOTE_PORT=$REMOTE_PORT \
    --env LOCAL_PORT=$REMOTE_PORT \
    --port $REMOTE_PORT \
    --restart=Never 


kubectl wait -n $NameSpace --for=condition=Ready pod/$TEMP_POD_NAME
kubectl port-forward -n $NameSpace pod/$TEMP_POD_NAME $LOCAL_PORT:$REMOTE_PORT

