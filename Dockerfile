FROM golang:1.20-buster
WORKDIR /app
ADD . .
RUN go mod tidy
RUN go build -o binapp
CMD ./binapp