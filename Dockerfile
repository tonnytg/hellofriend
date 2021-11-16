FROM golang
WORKDIR /app
COPY . .
RUN go build -o hello
CMD ["/app/hello"]
