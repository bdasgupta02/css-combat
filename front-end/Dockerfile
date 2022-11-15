FROM node:18.10.0-alpine

WORKDIR /front-end-app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./package*.json .

RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

ENV  WDS_SOCKET_PORT=3000
EXPOSE 3000

CMD ["npm", "start"]