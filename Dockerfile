FROM golang:1.25-alpine AS builder

ADD . /src
WORKDIR /src

RUN go build -o /app/edu_todo ./cmd/app/main.go

FROM alpine:3.18

COPY --from=builder /app/edu_todo /app/edu_todo

ENTRYPOINT [ "/app/edu_todo" ]




