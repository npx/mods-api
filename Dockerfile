FROM golang:1.10 as builder
WORKDIR /go/src/github.com/npx/mods-api
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/npx/mods-api/main .
EXPOSE 8000
CMD ["./main"]