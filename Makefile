# 定义变量
IMAGE_NAME = 172.25.10.67/test/sync
IMAGE_TAG = v1.0
DOCKERFILE = Dockerfile

# 构建 Docker 镜像
build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f $(DOCKERFILE) .

# 删除构建的 Docker 镜像
clean:
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG) || true

# 运行容器（假设容器的启动命令为 ./sync import produce_in）
run:
	docker run --rm $(IMAGE_NAME):$(IMAGE_TAG) ./sync import produce_in

# 列出所有 Docker 镜像
images:
	docker images

# 默认目标
.DEFAULT_GOAL := build
