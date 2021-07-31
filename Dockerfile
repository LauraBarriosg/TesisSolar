#imagen de golang
FROM golang:1.14.9-alpine as builder


LABEL maintainer="Laura Barrios <laura.barrios@correo.unimet.edu.ve>"

#para tarer cosas de git
RUN apk update \
 && apk add --no-cache git 
WORKDIR /go/src/app
COPY go.mod go.sum ./
COPY app.env ./
RUN go mod download 
COPY . .
RUN go build -o /go/bin/app

#mosquitto
#FROM eclipse-mosquitto:2.0.10 AS mqtt
#COPY /mosquitto/config/mosquitto.conf /mosquitto/config/mosquitto.conf




#Alpine
FROM alpine:latest
RUN apk add --no-cache openssl

#-------------------------------------------------- para esperar a los contenedores ----------------------------------
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
#ARG DB_CONNECTION_STRING
#ENV DB_CONNECTION_STRING=${DB_CONNECTION_STRING:-"host=postgres port=5432 user=EstacionSolar dbname=Tesis_Estacion password=Tesis sslmode=disable"}
COPY --from=builder /go/bin/app /go/bin/app 
COPY app.env .
COPY .env .
#Para correr el main
CMD ["/go/bin/app"]