FROM alpine/git

ENV TZ=Asia/Shanghai

WORKDIR /app

RUN git clone http://git.yungov.cn/RDC/code-explore.git


FROM golang:1.14.1-alpine

COPY --from=0 /app/code-explore /app/code-explore

WORKDIR /app/code-explore

# 如果国内网络不好，可添加以下环境
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN export GO111MODULE=on
# RUN export GOPROXY=https://goproxy.cn

RUN mkdir /opt/code-explore && ls /app/code-explore
RUN go build -o /opt/code-explore/code-explore ./ \
    && cp -r ./conf/ /opt/code-explore \
    && cp -r ./install/ /opt/code-explore\
    && cp ./scripts/run.sh /opt/code-explore\
    && cp -r ./static/ /opt/code-explore\
    && cp -r ./views/ /opt/code-explore\
    && cp -r ./logs/ /opt/code-explore\
    && cp -r ./docs/ /opt/code-explore
CMD ["/opt/code-explore/code-explore", "--conf", "/opt/code-explore/conf/default.conf"]