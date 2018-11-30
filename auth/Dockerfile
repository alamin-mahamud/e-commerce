FROM golang:alpine
WORKDIR /go/src/github.com/alamin-mahamud/e-commerce-go-api/auth

ENV AUTH_SRC_DIR=/go/src/github.com/alamin-mahamud/e-commerce-go-api/auth
ENV DEP_URL=https://raw.githubusercontent.com/golang/dep/master/install.sh
ENV BUILD_PACKAGES="git curl"

ADD . $AUTH_SRC_DIR

RUN	apk update                            \
	&& apk add --no-cache $BUILD_PACKAGES \
	&& curl $DEP_URL | sh                 \
	&& dep ensure                         \
	&& apk del $BUILD_PACKAGES

RUN cd $AUTH_SRC_DIR; \
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o main

ENTRYPOINT ["./main"]
