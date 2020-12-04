set -ex

containerId=$(docker ps |grep etcd |awk '{print $1}')
/usr/bin/docker exec -it $(containerId) /bin/sh -c 'etcdctl user add root --new-user-password 123456'
/usr/bin/docker exec -it $(containerId) /bin/sh -c 'etcdctl role add root'
/usr/bin/docker exec -it $(containerId) /bin/sh -c 'etcdctl user grant-role root root'
/usr/bin/docker exec -it $(containerId) /bin/sh -c 'etcdctl auth enable'

curl -ig 'localhost:2379/ping'
