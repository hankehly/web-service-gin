FROM golang:alpine AS build
WORKDIR /code
COPY . /code
RUN go build main.go
# Tried multi-stage build using scratch/alpine as runtime but kept encountering this error:
# standard_init_linux.go:228: exec user process caused: no such file or directory
FROM golang:alpine
COPY --from=build /code/main /go/bin/main
ENTRYPOINT [ "/go/bin/main" ]
