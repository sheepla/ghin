FROM golang:1.19-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o ghin .

FROM alpine:3.16 as runtime
RUN adduser -S -D -H -h /app ghin
USER ghin
COPY --from=builder /build/ghin /app/
ENTRYPOINT /app/ghin
