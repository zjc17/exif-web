FROM alpine:3.16
RUN rm -rf /var/cache/apk/*
WORKDIR /
COPY exif-web /bin/exif-web
ENTRYPOINT ["exif-web"]