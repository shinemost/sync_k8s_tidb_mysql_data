# 使用 Go 1.22 的官方基础镜像
FROM golang:1.22 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载 Go 依赖
RUN go mod download

# 复制所有的 .go 文件、config 目录和其他需要的目录到工作目录，排除 data 目录
COPY . .

# 构建 Go 应用程序
RUN CGO_ENABLED=0 go build -o sync .

# 使用更小的基础镜像来构建最终镜像
FROM scratch

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制构建好的应用程序
COPY --from=builder /app/sync .

# 复制 data 目录到最终镜像
COPY data/ data/

# 设定容器启动时的命令
CMD ["sync", "import", "produce_in"]
