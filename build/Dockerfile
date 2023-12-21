FROM golang:1.21.3 AS builder
# download modules
WORKDIR /src
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# build
FROM builder AS runner
ARG BUSINESS_NAME
ARG SERVICE_NAME
ARG SERVICE_TYPE

WORKDIR /douyin
COPY . .

RUN GOPROXY=https://goproxy.cn go build -o bin/${SERVICE_NAME}-${SERVICE_TYPE} /douyin/app/${BUSINESS_NAME}/${SERVICE_NAME}/${SERVICE_TYPE}/cmd/server

# run
FROM debian:stable-slim
ARG BUSINESS_NAME
ARG SERVICE_NAME
ARG SERVICE_TYPE

ENV SERVICE_NAME ${SERVICE_NAME}
ENV SERVICE_TYPE ${SERVICE_TYPE}

WORKDIR /app
COPY --from=runner /douyin/bin/${SERVICE_NAME}-${SERVICE_TYPE} bin/
COPY --from=runner /douyin/app/${BUSINESS_NAME}/${SERVICE_NAME}/${SERVICE_TYPE}/configs configs/

RUN if [ "$SERVICE_NAME" = "publish" ] && [ "$SERVICE_TYPE" = "job"]; then \
    apt-get update && \
    apt-get install --no-install-recommends -y ffmpeg ; \
fi
CMD bin/${SERVICE_NAME}-${SERVICE_TYPE} -conf configs