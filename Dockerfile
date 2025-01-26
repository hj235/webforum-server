FROM golang:1.20

# Set destination for copy
WORKDIR /app

# Prepare go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN go build -o ./serverbin ./cmd/server/main.go

# Bind to port 4000
EXPOSE 4000

# run
CMD [ "/app/serverbin" ]
