# Build stage
FROM golang:1.21.0 AS builder
ENV GOPROXY=https://proxy.golang.com.cn
#ENV GOPROXY=$GOPROXY

WORKDIR /app

COPY . .

#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/go-starter-gin-gorm ./cmd/app
RUN bash ./scripts/build.sh

# Production stage
FROM amd64/alpine:latest

# Set System TimeZone
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /app/bin/main .

ENV WORKDIR /app
CMD ./main start
