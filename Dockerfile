FROM golang:1.18.6

WORKDIR /bcc-app

# Copy the Pre-built binary file
COPY build/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
