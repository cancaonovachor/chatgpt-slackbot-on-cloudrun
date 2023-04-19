FROM golang:1.19-alpine3.17 as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./ ${ROOT}

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.17 as final
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app .
ENV PORT 8080
CMD ["./app"]
