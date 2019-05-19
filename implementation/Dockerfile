FROM golang:1.12.5 as base

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    apt update && \
    apt install -y netcat && \
    go get github.com/cespare/reflex

ENV CGO_ENABLED=0
WORKDIR ${GOPATH}/src/github.com/Angry-Potato/go-pay-me/implementation

COPY . .

RUN make deps



FROM base as test
ENTRYPOINT [ "make" ]
CMD ["test"]


FROM base as builder
RUN make build && mv app /app



FROM alpine:3.7 as prod
RUN apk --update add bash
COPY --from=builder /app /app
COPY wait-for.sh /wait-for.sh
ENTRYPOINT [ "/app" ]
CMD [ "" ]