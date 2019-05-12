FROM golang:1.12.5 as base

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    apt update && \
    apt install -y netcat

WORKDIR ${GOPATH}/src/github.com/Angry-Potato/go-pay-me/implementation

COPY . .

RUN dep ensure



FROM base as test
ENTRYPOINT [ "go" ]
CMD [ "test", "./..." ]



FROM base as builder
ENV CGO_ENABLED=0
RUN go build -o app && mv app /app



FROM alpine:3.7 as prod
RUN apk --update add bash
COPY --from=builder /app /app
COPY wait-for.sh /wait-for.sh
ENTRYPOINT [ "/app" ]
CMD [ "" ]