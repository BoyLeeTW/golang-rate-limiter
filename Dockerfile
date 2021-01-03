FROM golang:1.14-alpine
WORKDIR /rate-limiter
ADD . /rate-limiter
RUN cd /rate-limiter && go build ./cmd/rate-limiter
ENTRYPOINT [ "./rate-limiter" ]