FROM postgres:12.6-alpine
ENV LANG ja_JP.utf8

COPY ./docker/postgres/init/*.sql /docker-entrypoint-initdb.d/
