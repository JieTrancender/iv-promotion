set -ex

etcdctl user add root --new-user-password 123456
etcdctl role add root
etcdctl user grant-role root root
etcdctl auth enable

curl -ig 'localhost:2379/ping'
