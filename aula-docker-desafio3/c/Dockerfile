FROM golang:1.15

WORKDIR /golang/

COPY . .

RUN go mod download

RUN GOOS=linux go build c.go

EXPOSE 9092

ENTRYPOINT [ "./c" ]
