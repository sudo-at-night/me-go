FROM alpine:3.11.3

RUN adduser --disabled-password go
USER go
WORKDIR /home/go
COPY --chown=go ssr .

ENTRYPOINT ["./ssr"]
