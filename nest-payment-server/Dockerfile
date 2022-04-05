FROM node:12.13-alpine

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

RUN chmod +x wait-for-postgres.sh

CMD ["npm", "run", "start"]
