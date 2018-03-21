<h3 align="center">Guuid</h3>
<p align="center">Guuid 是一款由 Go 语言开发的 UUID 生成服务.</p>
<p align="center">
    <a href="https://travis-ci.org/dreamans/guuid"><img src="https://travis-ci.org/dreamans/guuid.svg?branch=master"></a>
</p>

---



UUID(Universally Unique Identifier)全局唯一标识符,是指在一台机器上生成的数字，它保证对在同一时空中的所有机器都是唯一的。

Guuid使用了服务器主机名、运行的进程ID、时间戳、随机数、时序元素等一系列元素来保证生成UUID的唯一性。

UUID 16 bytes, 构成: 
* 4 bytes 主机名&进程ID
* 4 bytes 时间戳
* 4 bytes 计数器
* 4 bytes 随机数

## 安装

```
$ go get -u github.com/dreamans/guuid/guuid-server
$ cd $GOPATH/src/github.com/dreamans/guuid/guuid-server
$ go install
```

## 运行
```
$GOPATH/bin/guuid-server -port 11223 -timeout 3
```
Usage:

    -port 运行端口

    -timeout 读写超时时间

运行成功:
```
⇨ http server started on [::]:11223
```

## RestAPI

* 获取UUID
```
$ curl http://127.0.0.1:11223/get

{
    "code": 200,
    "message": "OK",
    "uuid": "5d1b3abf-5ab2-2792-0000-00046b2155a4"
}
```

* 批量获取UUID
```
$ curl http://127.0.0.1:11223/mget/100

{
    "code": 200,
    "message": "OK",
    "data": [
        "5d1b3abf-5ab2-27cd-0000-00056d61a5e0",
        "5d1b3abf-5ab2-27cd-0000-000648f72a80",
        ...
    ]
}

说明:
最多返回1000条记录
```

* 获取简版UUID
```
$ curl http://127.0.0.1:11223/get/simple

{
    "code": 200,
    "message": "OK",
    "uuid": "5d1b3abf5ab22792000000046b2155a4"
}
```

* 批量获取简版UUID
```
$ curl http://127.0.0.1:11223/mget/100/simple

{
    "code": 200,
    "message": "OK",
    "data": [
        "5d1b3abf5ab227cd000000056d61a5e0",
        "5d1b3abf5ab227cd0000000648f72a80",
        ...
    ]
}
```
## 部署

LVS -> Nginx -> guuid-server

<img src="https://raw.githubusercontent.com/dreamans/guuid/master/docs/guuid-server.png">

* Configure Nginx
```
upstream guuid_server {
    server localhost:11211;
    server localhost:11212;
    server localhost:11213;
}

server {
    listen          80;
    server_name     localhost;

    location / {
        proxy_pass      http://guuid_server;
    }
}
```
