FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app/
ENV GOCACHE=/root/.cache/go-build
RUN go build

# Use Distroless image to reduce container image size
FROM gcr.io/distroless/base-debian12 as runner

WORKDIR /app

COPY --from=builder /app/gin-test /app/

EXPOSE 8080

CMD [ "./gin-test" ]