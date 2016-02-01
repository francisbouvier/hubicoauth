# HubiC oauth
# VERSION       0.1

FROM buildpack-deps:jessie

MAINTAINER Francis Bouvier, francis.bouvier@gmail.com

COPY hubicoauth /usr/local/bin

EXPOSE 8085

CMD ["hubicoauth"]
