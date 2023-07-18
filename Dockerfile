FROM golang:1.20-alpine as builder

WORKDIR /usr/src/app

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN mkdir bin

RUN go build -o /usr/bin/app ./main.go

FROM scratch:latest

WORKDIR /usr/bin

COPY --from=builder /usr/bin/app .
COPY --from=builder /usr/src/app/.env .
COPY --from=builder /usr/src/app/firebase-credentials.json .

EXPOSE 8000
CMD ["./app"]