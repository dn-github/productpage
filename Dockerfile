FROM alpine:3.4

COPY productpage .
ENTRYPOINT ./productpage
EXPOSE 3000