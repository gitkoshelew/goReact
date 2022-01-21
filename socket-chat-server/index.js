import express from "express";
import http from "http";
import { Server } from "socket.io";
import cors from "cors";
import dotenv from "dotenv";
import sequelize from "./db.js";
import router from "./src/routes/index.js";
import { Message } from "./src/models/models.js";

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

console.log(process.env.CHAT_SERVER_PORT);

const start = async () => {
  try {
    await sequelize.authenticate();
    await sequelize.sync();

    let connectedUsers = [];

    const addUser = (userId, socketId, receiverId) => {
      !connectedUsers.some((user) => user.userId === userId) &&
        connectedUsers.push({ userId, socketId, receiverId });
    };
    const removeUser = (socketId) => {
      connectedUsers = connectedUsers.filter(
        (user) => user.socketId !== socketId
      );
    };
    const getUser = (userId) => {
      return connectedUsers.find((user) => user.userId === userId);
    };

    io.on("connection", (socket) => {
      const { userId, receiverId } = socket.handshake.query;
      addUser(userId, socket.id, receiverId);
      console.log("Connected users", connectedUsers);

      socket.on(
        "USER_SEND_MESSAGE",
        async ({ senderId, receiverId, conversationId, text }) => {
          const message = await Message.create({
            senderId,
            receiverId,
            conversationId,
            text,
          });
          const targetUser = getUser(receiverId);
          io.to(socket.id).emit("SERVER_SEND_MESSAGE", message);
          if (targetUser && targetUser.receiverId === senderId) {
            io.to(targetUser.socketId).emit("SERVER_SEND_MESSAGE", message);
          }
        }
      );

      socket.on("disconnect", () => {
        removeUser(socket.id);
        console.log("Connected users after leave", connectedUsers);
      });
    });

    server.listen(PORT, () => console.log(`Server starts on port ${PORT}`));
  } catch (e) {
    console.log(e);
  }
};

start();
