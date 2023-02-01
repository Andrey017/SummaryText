FROM golang:latest
WORKDIR /api_gateway
ADD go.mod .
COPY ./ ./
RUN go build -o main ./cmd/main.go
#CMD [ "./main" ]