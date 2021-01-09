FROM golang:1.15-alpine as build

WORKDIR /app

RUN apk --no-cache add make ca-certificates

COPY . /app/

RUN make install_deps && make

FROM scratch

COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /app/build/app /usr/bin/app

ENTRYPOINT [ "/usr/bin/app" ]

EXPOSE 8080
