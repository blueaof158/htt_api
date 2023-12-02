FROM golang:1.21.4  as builder
# FROM golang:1.21.3-alpine:3.18 as builder
WORKDIR /app
COPY . /app

RUN go build -o main main.go


# # Build small image

FROM golang:1.21.4
# FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main . 
EXPOSE 8080

CMD [ "/app/main" ]