# syntax=docker/dockerfile:1

FROM node:20-alpine3.17

RUN mkdir /app

WORKDIR /app

COPY package.json .

RUN npm install --quiet

EXPOSE 5173

CMD npm run dev