import { Message } from "../../models/models.js";

export const startChat = (io) => {
  let connectedUsers = [];

  const addUser = (producerId, socketId, consumerId) => {
    !connectedUsers.some((user) => user.producerId === producerId) &&
      connectedUsers.push({ producerId, socketId, consumerId });
  };
  const removeUser = (socketId) => {
    connectedUsers = connectedUsers.filter(
      (user) => user.socketId !== socketId
    );
  };
  const getUser = (userId) => {
    return connectedUsers.find((user) => user.producerId === userId);
  };

  io.on("connection", (socket) => {
    console.log("Connected", socket.id);
    const { producerId, consumerId } = socket.handshake.query;
    addUser(producerId, socket.id, consumerId);
    console.log("Connected users", connectedUsers);

    socket.on(
      "USER_SEND_MESSAGE",
      async ({ producerId, consumerId, conversationId, text }) => {
        const message = await Message.create({
          producerId,
          consumerId,
          conversationId,
          text,
        });
        const targetUser = getUser(String(consumerId));
        io.to(socket.id).emit("SERVER_SEND_MESSAGE", message);
        if (targetUser && targetUser.consumerId === String(producerId)) {
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
