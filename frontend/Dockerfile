FROM node:alpine

RUN mkdir /app

WORKDIR /app

COPY ./package.json ./package.json
COPY ./package-lock.json ./package-lock.json

RUN npm i

COPY ./build ./build

CMD [ "npx", "serve", "-s", "build" ]