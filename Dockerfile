FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . ./
RUN swag init

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/root/.cache/go-build go build

# Use Distroless image to reduce container image size
FROM gcr.io/distroless/base-debian12 as runner

WORKDIR /app

COPY --from=builder /app/gin-test /app/

EXPOSE 8080

CMD [ "./gin-test" ]