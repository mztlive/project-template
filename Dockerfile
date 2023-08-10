# 使用官方Golang镜像作为构建环境
FROM golang:1.19 AS build-env

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"


# 设置工作目录
WORKDIR /app

# 拷贝项目文件到工作目录
COPY . .

# 使用Makefile构建项目
RUN make build

# 使用一个新的轻量级基础镜像
FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

# 创建部署目录并拷贝配置文件
WORKDIR /deploy
COPY deploy/ ./
RUN chmod +x ./run.sh && mkdir -p /deploy/log

# 从构建环境中拷贝编译出来的二进制文件（请根据Makefile的行为进行适当的修改）
COPY --from=build-env /app/bin/linux ./

EXPOSE 8849 11001

# 设置入口点为我们的应用（请根据你的应用实际的名字修改"output-binary-name"）
ENTRYPOINT ["/deploy/run.sh"]
