FROM alpine:3.2
ADD build/bin/foo /app/main
WORKDIR /app
ENTRYPOINT [ "/app/main" ]
