## Basic build
FROM golang:1.22-bookworm AS build
WORKDIR /code
COPY ./ ./
RUN go build -o /example ./main.go

## Create small dist image
FROM gcr.io/distroless/base-debian12 as prod
COPY --from=build /example /example
ENTRYPOINT ["./example"]