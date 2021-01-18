FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    GOOS=linux
# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o out/word-highlights-scraper-linux ./

# Build a small image
FROM golang:alpine

COPY --from=builder /build/out/word-highlights-scraper-linux /word-highlights-scraper

# Env variables for command line flags
ENV url=""
ENV depth=""

# Command to run
CMD /word-highlights-scraper -url=$url -depth=$depth
