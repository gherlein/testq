FROM scratch
MAINTAINER Greg Herlein <gherlein@herlein.com>
COPY ./testq /testq
ENTRYPOINT ["/testq"]

