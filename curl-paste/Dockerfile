FROM golang AS build
WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 go build -o curl-paste curl-paste.go config.go

FROM gcr.io/distroless/static-debian11
WORKDIR /
COPY --from=build /src/curl-paste /src/curl-paste.conf /
COPY --from=build /src/index.html /
ENTRYPOINT ["/curl-paste"]
