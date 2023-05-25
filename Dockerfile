FROM golang:alpine AS builder

# 在容器中设置工作目录
WORKDIR /app

# 下载项目需要的所有依赖
# 为了利用 Docker 的缓存层机制，我们应该先添加 go.mod 和 go.sum，然后再下载依赖。
# 这样只有在这两个文件发生变化时，Docker 才会重新下载依赖。
ADD go.mod go.sum ./
RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o laf-exporter .

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/laf-exporter .

CMD ["./laf-exporter"]
