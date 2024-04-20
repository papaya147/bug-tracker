FROM alpine:latest

RUN mkdir /app

COPY ./server-app /app/server-app
COPY ./db/migration /db/migration
COPY ./verification-email.html /verification-email.html

CMD [ "/app/server-app" ]