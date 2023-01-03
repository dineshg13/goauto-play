FROM golang:1.14.5-alpine3.12

ADD . /goplay
WORKDIR /goplay
RUN go mod download
RUN go build ./apps/hello-service


EXPOSE 6543

ENTRYPOINT ./hello-service
