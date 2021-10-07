FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /dist

EXPOSE 8080
CMD ["/app/dist"]