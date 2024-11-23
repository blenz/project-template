FROM golang:1.22-alpine AS base
WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download

FROM base as dev
RUN go install github.com/air-verse/air@latest
COPY . .
CMD [ "air", "--build.cmd", "go build -o /bin/app cmd/app/main.go", "--build.bin", "app"]

FROM base AS build
COPY . .
RUN go build -o app .

FROM alpine:3
COPY --from=build /bin/app .
CMD [ "app" ]
