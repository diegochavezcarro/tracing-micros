FROM golang:1.16-alpine as build

WORKDIR /micro
COPY . /micro/
RUN go build -o micro

FROM alpine:latest 
COPY --from=build micro /
CMD ./micro