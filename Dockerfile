# Using ubuntu 22.04, alpine v3.14 which still has wkhtmltopdf is broken
# docker build --rm -f Dockerfile -t go-genpdf:latest .
# docker run -d -p 8081:8080 --name go-genpdf-run --restart unless-stopped go-genpdf
FROM ubuntu:22.04 AS builder

# Add Git and other dependencies required for wkhtmltopdf
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates openssl git golang-go wkhtmltopdf xvfb libfontconfig1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from the Alpine base image
FROM ubuntu:22.04

WORKDIR /app

# Install wkhtmltopdf and its dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends wkhtmltopdf xvfb libfontconfig1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main /app/main

# Copy the templates folder
COPY --from=builder /app/templates /app/templates

# Expose the server port
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]

