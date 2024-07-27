FROM golang:1.21.4  as builder
# FROM golang:1.21.3-alpine:3.18 as builder
WORKDIR /app

COPY . /app

RUN go get 

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD [ "./main" ]