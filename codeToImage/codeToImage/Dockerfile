# Stage 1: Build the binary
FROM python:3.9-slim AS builder

# Set the working directory
WORKDIR /app

# Copy your application code into the container
COPY app.py /app/

# Install PyInstaller
RUN pip install pyinstaller

# Build the binary using PyInstaller
RUN pyinstaller --onefile app.py

# Stage 2: Runtime container
FROM debian:bullseye-slim

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/dist/app /app/

# Set executable permissions
RUN chmod +x /app/app

# Set the default command to execute the binary
CMD ["./app"]
