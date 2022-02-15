##
## Build
##

FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /huangcoolyserver

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /huangcoolyserver /huangcoolyserver

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/huangcoolyserver"]
