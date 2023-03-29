FROM golang:latest AS build
WORKDIR /src
COPY *.go go.* /src
RUN CGO_ENABLED=1 go build -o /bin/qdb

FROM scratch
WORKDIR /app
COPY --from=build /bin/qdb ./qdb
EXPOSE 3000
ENTRYPOINT ["/app/qdb"]