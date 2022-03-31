FROM node:12.13-alpine
WORKDIR /app

COPY package*.json ./

COPY ./ ./

RUN npm ci

RUN chmod +x wait-for-postgres.sh

CMD ["npm", "run", "start"]
