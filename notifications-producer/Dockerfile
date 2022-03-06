FROM node:16.14.0

WORKDIR /app

COPY . /app

RUN ls -la

RUN npm ci

RUN apt update
RUN apt install dos2unix
RUN dos2unix wait-for-postgres.sh
RUN chmod +x wait-for-postgres.sh
# RUN chmod +x wait-for-rabbit.sh

CMD ./wait-for-it.sh rabbitmq:5672 --timeout=0 --strict -- node --experimental-json-modules index.js
