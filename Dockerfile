FROM golang:1.18-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
#RUN apt-get update
#RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o cloud-test-task ./cmd/server/main.go

CMD ["./cloud-test-task"]