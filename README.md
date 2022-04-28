#使用grpc创建服务并注册etcd集群<br>
##etcd集群创建<br>
**参考: *https://www.cnblogs.com/wangtaobiu/p/16178759.html***
* 集群展示
  <br>| http://192.168.79.137:2379 | 95a497aa790a6957 |   3.5.1 |   25 kB |     false |      false |        13 |        138 |                138 |
  <br>| http://192.168.79.134:2379 | b7380bffd1c1af24 |   3.5.1 |   25 kB |      true |      false |        13 |        138 |                138 |
  <br>| http://192.168.79.136:2379 | c6a5db1cd31cb4ab |   3.5.1 |   25 kB |     false |      false |        13 |        138 |                138 |<br>
##grpc服务注册和服务发现<br>
**参考: *https://www.cnblogs.com/wangtaobiu/p/16173610.html***
* 服务注册
<br>func NewServiceDiscover() (Register, error)
* 服务发现
<br>func NewServiceDiscover() (Register, error)
