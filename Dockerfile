FROM golang:1.21-alpine

# Create app directory
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Use a non-root user to run the app
RUN adduser -D appuser
RUN chown -R appuser:appuser /app
USER appuser

CMD ["./main"]
EXPOSE 8081
