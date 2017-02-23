FROM debian:jessie-slim

RUN mkdir /plugins

ADD plug /plugins
ADD plugins/one/one.so /plugins
ADD plugins/two/two.so /plugins

CMD ["/plugins/plug", "/plugins/one.so", "/plugins/two.so"]