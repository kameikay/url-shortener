FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/app

FROM scratch
COPY --from=builder /app/server .
EXPOSE 3333
CMD [ "./server" ]