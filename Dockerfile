FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/github.com/crossphoton/iiitr-server/
WORKDIR /go/src/github.com/crossphoton/iiitr-server
RUN go mod download

COPY . /go/src/github.com/crossphoton/iiitr-server

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/iiitr-server github.com/crossphoton/iiitr-server


FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
ENV PORT=80
ENV DB_URL="postgres://postgres:postgres@database:5432/postgres?sslmode=disable"
COPY --from=builder /go/src/github.com/crossphoton/iiitr-server/build/iiitr-server /usr/bin/iiitr-server
EXPOSE 80

ENTRYPOINT ["/usr/bin/iiitr-server"]
