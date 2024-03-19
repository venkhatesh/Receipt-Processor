FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]