#
#    Dockerfile
#
#    Licensed under Creative Commons CC0.
#
#    To the extent possible under law, the author(s) have dedicated all copyright and related and
#    neighboring rights to this software to the public domain worldwide.
#    This software is distributed without any warranty.
#    You should have received a copy of the CC0 Public Domain Dedication along with this software.
#    If not, see <https:#creativecommons.org/publicdomain/zero/1.0/legalcode>.

# https://hub.docker.com/layers/golang/library/golang/1.14.15/images/sha256-6a39a02f74ffee82a169f2d836134236dc6f69e5946779da13deb3c7c6fedafd
FROM golang:golang@sha256:6a39a02f74ffee82a169f2d836134236dc6f69e5946779da13deb3c7c6fedafd

WORKDIR /tmp
EXPOSE 8000/tcp

ADD resources/favicon.ico /tmp
ADD pkg/src/eper.io/moose_audio/main.go /tmp
ADD resources/index.html /tmp
ADD resources/pcm.js /tmp

CMD go run main.go
