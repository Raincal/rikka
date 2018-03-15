FROM golang:1.10.0 as builder

LABEL maintainer="Raincal <cyj94228@gmail.com>"

ARG VCS_REF
ARG VCS_URL
ARG BUILD_DATE
ARG VERSION

LABEL org.label-schema.schema-version="1.0" \
    org.label-schema.version=$VERSION \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url=$VCS_URL \
    org.label-schema.vcs-type="Git" \
    org.label-schema.license="MIT" \
    org.label-schema.docker.dockerfile="/Dockerfile" \
    org.label-schema.name="Rikka"

WORKDIR $GOPATH/src/github.com/Raincal/rikka
COPY . $GOPATH/src/github.com/Raincal/rikka

RUN go get -v -d . && \
    go build -v . && \
    cp rikka $GOPATH/bin && \
    cp -R server $GOPATH/bin/ && \
    rm -rf $GOPATH/src

FROM alpine:3.7

# fix library dependencies
# otherwise golang binary may encounter 'not found' error
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /go/bin /

ENTRYPOINT ["/rikka"]