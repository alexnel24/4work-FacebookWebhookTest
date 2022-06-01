FROM golang:1.16-alpine

# and internal root ca certs
COPY .build/certs/*.crt /usr/local/share/ca-certificates/

# add the go app main to the container folder /app
COPY ./src/ /src/

# set the working directory in the container to /app
WORKDIR /src

# install ca certs, update certs, and create new directory /app
RUN apk add --no-cache ca-certificates bash curl \
RUN apk add --allow-untrusted --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/latest-stable/main/ ca-certificates \
  && update-ca-certificates \
  && apk add --no-cache bash curl \
  && apk --no-cache --available upgrade

WORKDIR /src

RUN go mod download

EXPOSE $PORT

CMD go run cmd/main.go