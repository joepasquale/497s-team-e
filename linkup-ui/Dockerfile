FROM node:alpine

WORKDIR /usr/app

ENV NODE_PATH /usr/app/node_modules

RUN npm i install shelljs --save

COPY package.json ./

COPY package-lock.json ./

COPY ./ ./

# add bash
RUN apk add --no-cache bash

RUN chmod +x ./env.sh

CMD ["npm", "start"]