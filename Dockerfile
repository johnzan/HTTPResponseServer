FROM golang:alpine as app-builder
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum main.go ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-response-server


FROM scratch
COPY --from=app-builder ./http-response-server ./http-response-server
EXPOSE 8080
# Run
CMD ["/http-response-server"]
