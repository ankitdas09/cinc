FROM alpine:latest as build-env

RUN apk update && \
    apk add build-base gcc musl-dev

FROM build-env

WORKDIR /app

# CMD ["gcc", "-o", "run","main.c"]
# docker run -it --rm -v ${PWD}/cmd/runner/main.c:/app/main.c runner sh -c "gcc main.c -o main && ./main"
# docker build -t runner ./cmd/runner