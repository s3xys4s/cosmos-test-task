FROM golang:1.18-alpine as build
WORKDIR /app
COPY . .
RUN go build

FROM alpine:latest
COPY --from=build /app/main .
CMD './main'