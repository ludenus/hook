FROM ubuntu:18.04

ADD pong /pong

ENV PONG_LISTENING_ADDRESS=":80"

ENTRYPOINT /pong
