#############################################
## Step 1: Compile the application for
##         the target architecture.
#############################################

FROM golang:1.19-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

# Install Git and make
RUN apk update
RUN apk add --no-cache git make

# Copy the source from the current directory to the Working Directory inside the container
WORKDIR /app
COPY . .

# Fetch the dependencies
RUN go get -d -v

# Build the binary
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} make build

#############################################
## Step 2: Create a smaller runner image
##         with only the compiled binary.
#############################################

FROM alpine

# Copy the pre-built binary file from the previous stage
WORKDIR /app
COPY --from=builder /app/bin/wosu /app/wosu

# Port 53 is used for DNS server
EXPOSE 53/udp

# Run the binary
ENTRYPOINT ["./wosu"]
