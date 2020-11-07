FROM golang:1.15

WORKDIR /golang/

COPY . .

RUN go mod download

RUN GOOS=linux go build b.go

EXPOSE 9091

ENTRYPOINT [ "./b" ]