FROM golang:latest
WORKDIR /article_service
ADD go.mod .
COPY ./ ./
RUN go build -o main ./cmd/main.go
CMD [ "./main" ]