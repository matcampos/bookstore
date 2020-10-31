
FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR /go/src/bookstore

ADD . /go/src/bookstore

RUN cp docker.env .env

RUN go get -d -v ./...

RUN go install bookstore

EXPOSE 3000

CMD ["bookstore"]