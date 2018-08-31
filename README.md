# monitor_kong_service
该项目的目的是为了方便监控kong上的upstream状态。

# 代码文件说明
- conf/ch.conf：      配置文件;请详细看下此文件。

# 前提
注：如果已有go环境，请忽略前两步。
1. 安装go环境  
请自行搜索安装go，版本不限；建议：go version go1.10.3 linux/amd64  
将go运行程序移动到/usr/bin/目录下  
执行go version验证go命令是否可用。 
2. 设置go环境
echo 'GOPATH=/data/go_pro/' > /etc/profile && source /etc/profile
3. 安装go依赖包 
go get github.com/astaxie/beego/config github.com/astaxie/beego/logs github.com/lib/pq 


# 运行
cd $GOPATH && \ 
git clone https://github.com/z843248880/monitor_kong_service.git && \ 
cd monitor_kong_service && \ 
go build . && \ 
./monitor_kong_service


