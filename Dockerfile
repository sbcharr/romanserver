FROM golang:1.9
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o server .
CMD ["/app/server"]
