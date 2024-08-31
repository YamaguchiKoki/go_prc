# Use official Golang image as the base
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download and install any needed Go modules
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Run the executable
CMD ["./main"]
