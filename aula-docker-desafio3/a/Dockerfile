FROM golang:1.15

WORKDIR /golang/

COPY . .

RUN go mod download

RUN GOOS=linux go build a.go

EXPOSE 9090

ENTRYPOINT [ "./a" ]