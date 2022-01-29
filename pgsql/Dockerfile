FROM postgres:latest

COPY db.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER=gopher
ENV POSTGRES_PASSWORD=gopherpw
ENV POSTGRES_DB=pubdb