import { Message } from "../models/models";

export const startSocket = (io) => {
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
};
