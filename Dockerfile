FROM golang:1.11.1-stretch

ADD entrypoint.sh /entrypoint.sh

RUN chmod 755 /entrypoint.sh

ENTRYPOINT /entrypoint.sh