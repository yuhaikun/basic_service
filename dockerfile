#------编译阶段------------------
# 设置基础镜像
FROM golang:1.20 AS builder


# 设置工作目录
WORKDIR /app

ENV GOPROXY="https://goproxy.cn,direct"
# 将依赖包复制到容器中
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中:
COPY . .

ENV CGO_ENABLED=0
# 构建可执行文件
RUN go build -o res


#------运行阶段-----------------
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 安装redis客户端依赖
RUN apk add --no-cache redis bind-tools

# 安装 tzdata 软件包
RUN apk --no-cache add tzdata

# 设置时区为CST
ENV TZ=Asia/Shanghai


# 从主机复制配置文件到容器中
COPY ./config/config.yaml /app/config/config.yaml

# 从构建阶段复制可执行文件到运行时阶段
COPY --from=builder /app/res .

# 暴露程序使用到端口号
#EXPOSE 6379
#EXPOSE 8080

# 在容器内安装 wget
RUN apk add --no-cache wget

# 获取容器的公网 IP 地址并写入环境变量
RUN PUBLIC_IP=$(wget -qO- http://ipecho.net/plain) && \
    echo "PUBLIC_IP=$PUBLIC_IP" >> /etc/environment
# 设置环境变量，指定redis地址
#ENV REDIS_ADDR=host.docker.internal:6379


RUN chmod +x res

# 运行可执行文件
CMD ["./res"]