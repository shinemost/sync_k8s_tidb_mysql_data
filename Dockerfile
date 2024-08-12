FROM alpine
# 设置工作目录
WORKDIR /root/

# 复制 data 目录到最终镜像
COPY data/ data/
COPY sync config.yaml /root/

RUN chmod +x sync
RUN ls -l /root/


