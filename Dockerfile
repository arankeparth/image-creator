# Stage 1: Build the Go application
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install dependencies, including GCC
RUN apt update && apt install -y gcc

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies (this layer is cached if go.mod and go.sum are not changed)
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create the final, minimal image
FROM debian:bookworm-slim

# Install necessary dependencies to run the Go binary (e.g., libc6)
RUN apt-get update && apt-get install -y libc6

# Set the Current Working Directory inside the container
WORKDIR /root/



# Copy the compiled Go binary from the builder stage

COPY --from=builder /app/main .
COPY --from=builder /app/codeToImage/fonts /root/codeToImage/fonts
COPY --from=builder /app/codeToImage/EditProfile.webp /root/codeToImage/EditProfile.webp
COPY --from=builder /app/codeToImage/CreateProfile.webp /root/codeToImage/CreateProfile.webp

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
