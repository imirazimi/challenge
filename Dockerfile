# Build Stage
FROM golang:1.23.0-alpine AS builder

# Set environment variable
ENV APP_NAME challenge

# Add a work directory
WORKDIR /$APP_NAME

# Copy app files
COPY . .

# Budild application
RUN CGO_ENABLED=0 go build -mod=vendor -v -o $APP_NAME .

# Run Stage
FROM alpine:3.20 AS runtime

# Copy the binary from the builder stage
COPY --from=builder /challenge/challenge .

# Expose application port
EXPOSE 3001

# Start the application
CMD ["/challenge"]
