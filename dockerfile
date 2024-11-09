FROM golang:1.22 AS builder

WORKDIR /builder

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apt-get update && apt-get install -y libmupdf-dev

COPY . .

RUN go build -o /main .

FROM debian:bullseye-slim

EXPOSE 5001

WORKDIR /

COPY --from=builder ./main .

COPY .env .

CMD ["/main"]
