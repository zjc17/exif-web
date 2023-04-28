# Exif Web

[Features](#features)
•
[Download](#download)
•
[Usage](#usage)
•
[Future](#future)
•
[Credit](#credit)

Lightweight opensource Exif Analysis Backend written in Golang, with binary size ~ 20MB.

Support Restfull API, WebUI, x86, ARM, Linux, macOS.

## Features

- 🏎️ Utilized fastest and wildly used EXIF Javascript [lib](https://github.com/MikeKovarik/exifr).
- 📷 Files: .jpg, .tif, .png, .heic, .avif, .iiq
- 📑 Reads only first few bytes for a given url or image data.
- 🗜️ Easy to deploy: one cross-platform binary file or docker to deploy.

## Download

Download the binaries for your system and architecture from the [releases page](https://github.com/zjc17/exif-web/releases).

If you use docker, you can use the following command ([Dockerhub](https://hub.docker.com/r/lovecho/exif-web))

```bash
docker pull lovecho/exif-web:latest
```

## Usage

Use default parameters launch the api server:

```bash
./exif-web
```

### Restful API Usage

Parse a remote image with its uri

```bash
curl 'http://127.0.0.1:8080/api/v1/parse?url=$IMAGE_URL'
```

### Docker Usage
There is no difference between using parameters in Docker and the above, 
for example, we start a Web UI formatting tool service in Docker:

```bash
docker run --rm -it -p 8080:8080 lovecho/exif-web:latest
```

## Future

- [ ] Support read image on local filesystem
- [ ] a simple web ui as live demo
- [ ] a simple build in k/v caching system to prevent duplication of parsing

## Credit

Exif parse components:
- [exifr](https://github.com/MikeKovarik/exifr): The fastest and most versatile JavaScript EXIF reading library, under [MIT license]
- modified javascript version for golang execution, under [Apache-2.0 license], 28/04/2023:
  - allow running in golang
  - https://github.com/zjc17/exif-web
Runtime dependent components:
- [goja](https://github.com/dop251/goja): ECMAScript 5.1(+) implementation in Go, under [MIT license].
Web components:
- [Gin](https://github.com/gin-gonic/gin): a HTTP web framework written in Go (Golang), under [MIT license]
