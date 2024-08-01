FROM golang:1.22 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the Go source file into the container
COPY . .

RUN apt-get update &&\
    apt-get install -y --no-install-recommends ca-certificates curl \
    tcpdump wget telnet iputils-ping dnsutils net-tools iptables

# Build the Go program inside the container
RUN go build -o server main.go

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 10014 \
  "choreo"
# Use the above created unprivileged user
USER 10014

# Set the entrypoint to the executable
ENTRYPOINT ["./server"]
