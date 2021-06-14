FROM golang:latest

WORKDIR /go/src/cartesianAPI

ADD . /go/src/cartesianAPI

RUN go build -o main . 

EXPOSE 8000

ENTRYPOINT [ "/go/src/cartesianAPI/main" ]