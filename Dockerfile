FROM golang:latest AS builder
WORKDIR /go_sample_login_register/
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM alpine
COPY --from=builder /go_sample_login_register/main .
CMD ["./main"]