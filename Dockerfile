FROM golang:alpine as app-builder
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum main.go ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
#COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /http-response-server


FROM scratch
# the test program:
COPY --from=app-builder ./http-response-server ./http-response-server
EXPOSE 8080
# Run
CMD ["/http-response-server"]