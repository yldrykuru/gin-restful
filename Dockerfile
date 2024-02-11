# Use the latest official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire local project directory into the container at /app
COPY . .

# Build the Go application and create an executable named 'main'
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Specify the command to run your application when the container starts
CMD ["./main"]
