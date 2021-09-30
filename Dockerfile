FROM golang:1.16.4 AS builder
COPY . /src
WORKDIR /src
ENV GOPROXY https://goproxy.cn
ENV CGO_ENABLED 0
RUN go mod download
RUN go build -o /dist/api ./cmd/api
RUN go build -o /dist/sync ./cmd/syncer

FROM juxuny/alpine:3.13.5 AS final
COPY --from=builder /dist /dist
ENTRYPOINT /dist/api