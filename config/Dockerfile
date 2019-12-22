FROM golang:latest
LABEL maintainer="Marco Barroso"
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get
RUN go install
ENTRYPOINT go run main.go
EXPOSE 8000