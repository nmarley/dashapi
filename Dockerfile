# build the project first, then copy only the binary
FROM golang:alpine
WORKDIR /build
RUN apk update && apk add --no-cache git
COPY . /build
RUN (cd /build && go get -d -v ./... && CGO_ENABLED=0 GOOS=linux go build)

# this results in an image the size of the binary (~10 MB)
FROM scratch
COPY --from=0 /build/dashapi /dashapi
ENTRYPOINT ["/dashapi"]

LABEL maintainer="Nathan Marley <nathan.marley@gmail.com>"
LABEL description="Dash Governance Data API"
