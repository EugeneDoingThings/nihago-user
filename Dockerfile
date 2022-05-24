# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN cd cmd && go build -o /nihago-user

EXPOSE 8090

CMD [ "/nihago-user" ]