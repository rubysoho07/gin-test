FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . /app/

RUN go get
RUN go build

FROM golang:1.21-alpine as runner

WORKDIR /app

COPY --from=builder /app/gin-test /app/

EXPOSE 8080

CMD [ "./gin-test" ]