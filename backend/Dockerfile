FROM golang:1.23-alpine AS base
WORKDIR /backend
COPY go.mod go.sum ./
RUN go mod download

FROM base as dev
RUN go install github.com/air-verse/air@v1.61.1
COPY . .
CMD ["air", "--build.cmd", "go build -o /bin/app cmd/app/main.go", "--build.bin", "app"]

FROM base AS build
COPY . .
RUN go build -o /bin/app .

FROM alpine:3
COPY --from=build /bin/app .
CMD [ "app" ]
