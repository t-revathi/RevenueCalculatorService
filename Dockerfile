FROM golang:1.16-alpine as build

WORKDIR /app

RUN apk update

RUN apk add gcc

RUN apk add g++

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

WORKDIR /app/cmd/userservice

RUN go build -o /docker-userservice

FROM golang:1.16-alpine

WORKDIR /app

COPY --from=build /docker-userservice /docker-userservice


EXPOSE 3333

CMD [ "/docker-userservice" ]