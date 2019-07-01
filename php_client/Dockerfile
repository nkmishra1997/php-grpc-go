FROM php:7.1-apache

LABEL maintainer="Nikhil Mishra <nikhil.mishra@zomato.com>"

WORKDIR .

COPY . .

# RUN apt-get update && apt-get install phpunit php-pear
RUN apt-get update && apt-get install -y zlib1g-dev

# install composer
RUN cd /usr/local/bin && curl -sS https://getcomposer.org/installer | php
RUN cd /usr/local/bin && mv composer.phar composer

RUN pecl install grpc

#install protoc
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
RUN mkdir -p /tmp/protoc && \
    curl -L https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip > /tmp/protoc/protoc.zip && \
    cd /tmp/protoc && \
    unzip protoc.zip && \
    cp /tmp/protoc/bin/protoc /usr/local/bin && \
    cd /tmp && \
    rm -r /tmp/protoc && \
    docker-php-ext-enable grpc

CMD [ "php", "./index.php" ]