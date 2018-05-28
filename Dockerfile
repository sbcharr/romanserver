FROM golang:1.9.6
RUN mkdir -p /go/src/github.com/squeakysimple/romanserver
WORKDIR /go/src/github.com/squeakysimple/romanserver
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["romanserver"]
