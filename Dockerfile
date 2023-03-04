FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -v -o server

FROM alpine
COPY --from=builder /app/server /server
COPY --from=builder /app/conf /conf
# 声明服务端口(不起实际作用，只是一个标识，具体启动端口是容器中程序的配置)
EXPOSE 8000
# docker build -t go_rent .
# docker run -itd -p 8001:8000 -v /Users/donghao/GolandProjects/rent_backend/conf/dev.env:/conf/dev.env --name=go_rent go_rent
# docker run -itd  --name=alpine alpine
