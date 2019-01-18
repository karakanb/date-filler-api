# Build the binary.
FROM golang:alpine AS builder
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Run the built binary on scratch image.
FROM scratch
COPY --from=builder /build/main /
ENV PORT 80
CMD ["./main"]