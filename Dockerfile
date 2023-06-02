FROM node as nodebuilder

ADD client /app
WORKDIR /app
RUN npm install
RUN npm run build

FROM golang as builder
WORKDIR /go/src/github.com/thomasfady/xsstower
COPY . .
ENV GIN_MODE release
COPY --from=nodebuilder /app/dist /go/src/github.com/thomasfady/xsstower/app
RUN ls -la /go/src/github.com/thomasfady/xsstower/app
RUN set -x && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o xsstower -ldflags="-s -w" main.go

FROM scratch
LABEL maintainer "Thomas FADY <thomas.fady@gmail.com>"

WORKDIR /root/
COPY --from=builder /go/src/github.com/thomasfady/xsstower/xsstower .
ENV GIN_MODE release
CMD ["./xsstower"]