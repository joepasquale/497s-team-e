FROM node:alpine

WORKDIR /usr/app

ENV NODE_PATH /usr/app/node_modules

RUN npm install shelljs
RUN npm install react-google-login

COPY package.json ./

COPY package-lock.json ./

COPY ./ ./

# add bash
RUN apk add --no-cache bash

RUN chmod +x ./env.sh

CMD ["npm", "start"]