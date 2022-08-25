set -e

# args
KubeCfg=uat.yml # 目标集群的kubeconfig -- kubecm的列表项
NameSpace=uat-eureka # 目标ns
REMOTE_HOST=35.239.84.150
REMOTE_PORT=5432 # 需要转接的目标port
LOCAL_PORT=5430 # 本地port
TEMP_POD_NAME=test-jump-redash-uat # 用于转接的pod名称

export KUBECONFIG=~/dev/repos/kubeconfig/$KubeCfg


# 脚本退出时自动清理掉pod
function cleanup {
  echo "cleanup..." 
  kubectl delete -n $NameSpace pod/$TEMP_POD_NAME --grace-period 1 --wait=false
}
trap cleanup EXIT

kubectl run -n $NameSpace --image marcnuri/port-forward $TEMP_POD_NAME \
    --env REMOTE_HOST=$REMOTE_HOST \
    --env REMOTE_PORT=$REMOTE_PORT \
    --env LOCAL_PORT=$REMOTE_PORT \
    --port $REMOTE_PORT \
    --restart=Never 


kubectl wait -n $NameSpace --for=condition=Ready pod/$TEMP_POD_NAME
kubectl port-forward -n $NameSpace pod/$TEMP_POD_NAME $LOCAL_PORT:$REMOTE_PORT &

# keep alive. ref:https://stackoverflow.com/questions/47484312/kubectl-port-forwarding-timeout-issue
while true ; do sleep 60 ; nc -vz 127.0.0.1 $LOCAL_PORT ; done 