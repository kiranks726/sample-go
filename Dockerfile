FROM ubuntu:latest

RUN apt-get update \
    && apt-get install -y software-properties-common

ENV PORT=5000

EXPOSE ${PORT}

WORKDIR /code

COPY . ./
