FROM alpine:3.5

COPY bin/rabbit /rabbit

EXPOSE 8888

ENTRYPOINT ["/rabbit"]

