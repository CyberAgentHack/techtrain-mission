FROM golang:1.15 as builder
WORKDIR /go/src/github.com/task4233/techtrain-mission/

ENV GO111MODULE="on"
COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o server

FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/task4233/techtrain-mission/server /server

EXPOSE 8080
CMD ["/server"]