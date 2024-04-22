# Golang image as a base
FROM golang:1.22


# metadata labels
LABEL maintainer="Elhadj Abdoul Diallo <elhhabdouldiallo@gmail.com>"
LABEL description="Dockerfile for Asci Art web application"
LABEL version="1.0"

# Set the Current Working Directory inside the container
WORKDIR /app

# Environment variables
ENV APP_DIR=app/web_app/main.go
ENV APP_BINARY=main
ENV PORT=8080

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o $APP_BINARY $APP_DIR

# Expose the port
EXPOSE $PORT

# Command to run the executable
CMD ["./main"]

# command to build docker image
# docker build --build-arg APP_DIR=app/web_app --build-arg APP_BINARY=main --build-arg PORT=8080 -t my-golang-app .

