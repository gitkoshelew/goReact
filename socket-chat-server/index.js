import express from "express";
import http from "http";
import { Server } from "socket.io";
import cors from "cors";
import dotenv from "dotenv";
import sequelize from "./db.js";
import router from "./src/routes/index.js";
import { startSocket } from "./src/services/socket.js";

dotenv.config();

const app = express();
const server = http.createServer(app);
const io = new Server(server, {
  cors: {
    origin: "http://localhost:3000",
  },
});

app.use(
  cors({
    origin: "http://localhost:3000",
    credential: true,
  })
);
app.use(express.json());

const PORT = process.env.CHAT_SERVER_PORT || 5000;

app.use("/api", router);

const start = async () => {
  try {
    await sequelize.authenticate();
    await sequelize.sync();

    startSocket(io);

    server.listen(PORT, () => console.log(`Server starts on port ${PORT}`));
  } catch (e) {
    console.log(e);
  }
};

start();
