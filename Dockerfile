FROM golang:1.24rc3-alpine as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download

# Copy the source from the current directory to the working directory in the container.
COPY . .

RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]