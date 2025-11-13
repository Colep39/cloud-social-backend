FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o auth ./cmd/auth

FROM gcr.io/distroless/base-debian12
COPY --from=build /app/auth /auth
EXPOSE 8081
USER nonroot:nonroot
ENTRYPOINT ["/auth"]
