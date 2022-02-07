FROM node:14.17.3

WORKDIR /app

COPY . /app

RUN ls -la

RUN npm install

RUN chmod +x wait-for-postgres.sh
# RUN chmod +x wait-for-rabbit.sh

CMD ./wait-for-it.sh rabbitmq:5672 --timeout=0 --strict -- node index.js