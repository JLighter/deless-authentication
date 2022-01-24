FROM golang AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S app && adduser -S app -G app
USER app
WORKDIR /home/app
COPY --from=builder /bin/app ./
EXPOSE 80
CMD ./app | grep "^{"
