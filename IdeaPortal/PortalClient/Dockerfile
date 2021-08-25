# Alpine is lightweight image
FROM golang:1.16-alpine as builder
WORKDIR /app
# Copy files to download modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy all other files and build .exe files
COPY . ./
RUN go build -o /PortalClient
EXPOSE 8000
# runs .exe file which was generated using go build
CMD ["/PortalClient"]