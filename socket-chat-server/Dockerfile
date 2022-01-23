FROM node:14.17.3

WORKDIR /app

COPY . /app

RUN npm install

#make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

CMD ["node", "index.js"]