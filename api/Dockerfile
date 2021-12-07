FROM golang:1.17

RUN go version
ENV GOPATH=/

# Set destination for COPY
WORKDIR /api

COPY ./ ./

#instal psql
RUN apt-get update
RUN apt-get -y install postgresql-client

#make witt-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o goreact ./main.go

CMD [ "/goreact" ]


