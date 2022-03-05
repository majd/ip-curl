FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build

EXPOSE 3000

CMD [ "/app/ip-curl" ]