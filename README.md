# grpc-go-helloworld
	这是一个grpc-go实例项目。
# 实验步骤
## 添加%GOPATH%/bin到系统环境变量
## 安装和配置protobuf
    在https://link.jianshu.com/?t=https://github.com/google/protobuf/releases/tag/v3.3.0下载protoc-3.3.0-win32.zip的包，获取其protoc.exe放置在%GOPATH%/bin
## 安装 protoc-gen-go
    1.自动编译获取protoc-gen-go.exe：在终端直接执行export GOPATH=/d/project/go/grpc-go-helloworld && go get -u github.com/golang/protobuf/protoc-gen-go，可以在你的%GOPATH%/bin路径下找到一个自动生成的protoc-gen-go.exe（这一步不能完成的话，考虑开个全局的翻墙）
    2.自己编译获protoc-gen-go.exe：取进入 https://github.com/golang/protobuf 的github库，通过“clone or download”获取整个protobuf项目；然后将其原名protobuf-master改名为protobuf放入 %GOPATH%/src/github.com/golang 下，进入protobuf的子项目protoc-gen-go，执行 go build &&go install ，最后确保%GOPATH%/bin已在环境变量下
## 安装grpc-go及text，net包
    （1） 依赖
    需要至少 go 1.7+，以及 gRPC

    （2）安装grpc-go

        需要注意的是：这里安装的 grpc-go 只是一个第三方库，提供的是库函数接口，而不是二进制程序。因此，grpc-go 安装之后提供的是第三方包的作用。
        1.自动安装：export GOPATH=/d/project/go/grpc-go-helloworld && go get -u google.golang.org/grpc
        2.手动下载：进入 https://github.com/grpc/grpc-go 通过“clone or download”获取整个grpc-go项目；将其改名为grpc放入 %GOPATH%/src/google.golang.org
    （3）安装其他包
        net,text，genproto包安装：
            net，text放置在 %GOPATH%/src/golang.org/x
            go-genproto更名爲genproto放置在 %GOPATH%/google.golang.org
            通过“clone or download” 获取:
            https://github.com/golang/text.git
            https://github.com/golang/net.git
            https://github.com/google/go-genproto
## 进行proto文件编写
    见 %GOPATH%/src/pb/helloworld.proto
## 利用protoc编译 .proto 文件
    在 %GOPATH%/src/pb 执行  protoc ./helloworld.proto --go_out=plugins=grpc:.
## 编写server和client
    见 %GOPATH%/src/client/main.go
    %GOPATH%/src/server/main.go
## 运行测试
	运行server和client