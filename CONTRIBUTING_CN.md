时序数据库主从同步系统构建
========================

使用源码构建**时序数据库主从同步系统**。


系统要求
--------

Linux 64位系统即可。


依赖环境
--------

本系统使用InfluxDB v1.5.2，所以需要***Go 1.9.2***或者以上的版本。

InfluxDB使用Dep管理依赖包，需要安装[dep](https://github.com/golang/dep)。

主从同步脚本使用***Python 2.7***开发。

需要从Github上拉取源码，环境中需要安装***git***。


安装InfluxDB
------------

新建目录`$YOUR_PATH/src/influxdb`，`$YOUR_PATH`为自定义目录：
```
mkdir -p $YOUR_PATH/src/influxdb
```

将源码下载解压到目录`$YOUR_PATH/src/influxdb`中：
```
cd $YOUR_PATH/src/
git clone $THIS_PATH
```

将目录`$YOUR_PATH/`设置为`GOPATH`：
```
export GOPATH=$YOUR_PATH/
```

进入目录中：
```
cd $YOUR_PATH/src/influxdb
```

安装依赖：
```
dep sure
```

构建安装二进制执行码：
```
go clean ./...
go install ./...
```

安装完成后，二进制执行码放在`$YOUR_PATH/bin/`中，启动时序数据库：
```
$YOUR_PATH/bin/influxd
```

安装主从同步脚本
----------------

新建目录`$YOUR_SCRIPT`:
```
mkdir $YOUR_SCRIPT
```

从Github中拉取脚本：
```
cd $YOUR_SCRIPT
git clone https://github.com/callELPSYCONGROO/master_slave
```

参考其中的[README.md](https://github.com/callELPSYCONGROO/master_slave/blob/master/README.md)完成脚本配置和运行。