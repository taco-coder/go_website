FROM golang:1.18-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o creds-go-service . 

EXPOSE 8080

CMD ["./creds-go-service"] 