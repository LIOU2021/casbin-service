FROM golang:1.21.3 AS build-stage

WORKDIR /app/src

COPY ./data/src/. .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /casbin-service

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app/src

COPY --from=build-stage /casbin-service /app/src/casbin-service

EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT ["/app/src/casbin-service"]