FROM golang:1.12.5 as base

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR ${GOPATH}/src/github.com/Angry-Potato/go-pay-me/implementation

COPY . .

RUN dep ensure



FROM base as builder
ENV CGO_ENABLED=0
RUN go build -o app && mv app /app



FROM scratch as prod
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]