# syntax=docker/dockerfile:1
FROM golang:1.17

RUN go version
ENV GOPATH=/

# Set destination for COPY
WORKDIR /api

#instal psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

#make witt-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go build -o goreact ./main.go

CMD [ "/goreact" ]


