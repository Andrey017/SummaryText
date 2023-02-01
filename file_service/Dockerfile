FROM golang:latest
WORKDIR /file_service
ADD go.mod .
COPY ./ ./
RUN go build -o main ./cmd/main.go
CMD [ "./main" ]