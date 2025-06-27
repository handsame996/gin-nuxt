FROM golang:alpine as builder

WORKDIR /go/src/gin-nuxt/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server ./cmd/app

LABEL MAINTAINER="1377231006@qq.com"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/gin-nuxt/server

COPY --from=0 /go/src/gin-nuxt/server/server ./
COPY --from=0 /go/src/gin-nuxt/server/configs ./configs/
COPY --from=0 /go/src/gin-nuxt/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml