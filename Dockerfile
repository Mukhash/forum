FROM golang:1.16
LABEL name="Forum"
LABEL description="Web Forum"
LABEL authors="Mukhash, AlmasChamp"
LABEL version="1.0"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8080