# Alpine is lightweight image
FROM golang:1.16-alpine as builder
WORKDIR /app
# Copy files to download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy all other files and build .exe 
COPY . ./
RUN go build -o /app/PortalClient

# Use multi-stage to reduce docker image size
FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8000
CMD ["/app/PortalClient"]