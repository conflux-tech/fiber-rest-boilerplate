FROM golang:1.14 as build

WORKDIR /go/src/build
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o ./app ./main.go

FROM alpine:latest

WORKDIR /go/src/deploy
COPY --from=build /go/src/build .
RUN chmod +x ./app/main

EXPOSE 3000

CMD ["./app/main"]
