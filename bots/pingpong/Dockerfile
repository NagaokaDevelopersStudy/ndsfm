FROM alpine:3.8

RUN apk --update --no-cache add ca-certificates

ADD pingpong.alpine /pingpong
ADD pingpong.json /pingpong.json
ADD entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
