FROM golang:1.22.1 as builder

LABEL maintainer="Oleksandr Goltsman"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-store-things .

# 

FROM alpine:3.19.1

WORKDIR /root/

COPY --from=builder /app/go-store-things .

ENTRYPOINT ["./go-store-things"] 
