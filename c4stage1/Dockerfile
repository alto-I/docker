From golang:1.13.4-alpine3.10
WORKDIR /src
COPY ./main.go /src
RUN go build -o /usr/local/bin/startapp main.go
WORKDIR /
CMD [ "/usr/local/bin/startapp" ]
