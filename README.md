# gmf
go语言微服务框架  
技术特点：  
一、一次编译多服部署，告别传统微服务框架每个微服务都需要编译一次的麻烦，使用指令启动微服务  
二、低耦合，虽然只需要编译一次，但是微服务之间的耦合度非常低  
三、易扩展，微服务采用接口设计，实现接口即可实现一个微服务  
四、易使用，提供命令行工具，一条指令创建一个微服务  
快速开始：  
go版本1.16.2  
#一、下载gmf框架  
git clone https://github.com/lijianjunljj/gmf.git  
go mod tidy //下载依赖包  
go build -o gmf.exe main.go //编译  
二、启动etcd  
windows版本etcd快速下载地址：链接: https://pan.baidu.com/s/1RO6kwp8WNWhKga8J6PCYoA?pwd=gmf1 提取码: gmf1  
windows版本etcdkeeper快速下载地址：链接: https://pan.baidu.com/s/1KK0UT5Itd4ijJKCjlRhNGg?pwd=gmf1 提取码: gmf1  
运行etcd.exe  
运行etcdkeeper.exe 输入http://127.0.0.1:8080/etcdkeeper/可以查看已经启动和注册的微服务  
三、启动网关  
windows:.\gmf.exe start --name gateway  --config .\config\development.yaml  
linux:  
chmod +x gmf  
./gmf start --name gateway  --config ./config/development.yaml  
四、创建mysql数据库gmf  
五、启动用户服务  
windows:.\gmf.exe start --name  userServer  --config .\config\development.yaml  
linux:./gmf start --name userServer  --config ./config/development.yaml  
六、测试是否成功  
http://127.0.0.1:4000/ping 返回pong，说明安装成功  
快速开发：  
一、创建微服务：\gmf.exe create my_service输入此命令后会自动创建微服务的代码  
二、修改模型：my_service/model  
三、编写go-micro的probuf协议文件.proto放到my_service/services/protos，编译出来的协议文件放到my_service/services,
不会使用probuf的请参考此教程：https://www.cnblogs.com/qi66/p/17034076.html  
四、自定义路由：my_service/router  
五、自定义路由回调：my_service/handler  
六、修改核心逻辑：my_service/core  
七、修改熔断器：my_service/wrappers  