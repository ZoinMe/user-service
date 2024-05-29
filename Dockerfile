# Use the base image with Go installed
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o ./out/zoinme-user-services .

# Final stage: Use a minimal base image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage into the current stage
COPY --from=builder /app/out/zoinme-user-service .

# Expose port 8080 for the application
EXPOSE 8080

# Define the command to run the application
CMD ["./zoinme-user-service"]
