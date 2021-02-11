FROM golang:1.15-alpine AS build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY src/ .
RUN go build main.go


FROM scratch
COPY --from=build /build/main /app/main
CMD ["/app/main"]

