FROM golang:1.17-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk add git make
WORKDIR /go/observe_my_hatebu

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o app

FROM alpine
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /go/observe_my_hatebu/app /app
CMD ["/app"]
