import amqp from "amqplib";
import rabbitDefinitions from "../../../../rabbit/definitions.json";

const rabbitSettings = {
  protocol: "amqp",
  // hostname: process.env.RABBIT_HOST,
  hostname: "localhost",
  // port: process.env.RABBIT_PORT,
  port: 5673,
  username: process.env.RABBIT_USER,
  password: process.env.RABBIT_PASSWORD,
  vhost: "/",
  authMechanism: ["PLAIN", "AMQPLAIN", "EXTERNAL"],
};

const NOTIFICATIONS_EXCHANGE = rabbitDefinitions.exchanges.find(
  (exchange) => exchange.name === "NOTIFICATIONS_EXCHANGE"
).name;

const BROKER_RECEIVED_NOTIFICATION = "BROKER_RECEIVED_NOTIFICATION";
const CLIENT_RECEIVED_NOTIFICATION = "CLIENT_RECEIVED_NOTIFICATION";

export const startNotificationConsumer = async (io) => {
  try {
    const connection = await amqp.connect(rabbitSettings);
    const channel = await connection.createChannel();

    await channel.assertExchange(NOTIFICATIONS_EXCHANGE, "direct", {
      durable: true,
    });

    const notificationQueues = rabbitDefinitions.queues.filter(({ name }) => {
      const queuePrefix = name.slice(0, name.indexOf("_"));
      return queuePrefix === "NOTIFICATIONS";
    });

    const assertQueuesArray = [];

    notificationQueues.forEach((queue) => {
      assertQueuesArray.push(
        async () =>
          await channel.assertQueue(queue, { durable: true, autoDelete: false })
      );
    });

    await Promise.all(assertQueuesArray);

    const bindings = rabbitDefinitions.bindings.filter(({ destination }) => {
      const destinationPrefix = destination.slice(0, destination.indexOf("_"));
      return destinationPrefix === "NOTIFICATIONS";
    });

    bindings.forEach(async ({ destination, source, routing_key }) => {
      await channel.bindQueue(destination, source, routing_key);
    });

    const connectedUsers = new Map();

    function getUserToRemove(map, searchValue) {
      for (let [key, value] of map.entries()) {
        if (value === searchValue) return key;
      }
    }

    const addUser = (userId, socketId) => {
      if (!connectedUsers.has(userId)) {
        connectedUsers.set(userId, socketId);
      }
    };
    const removeUser = (socketId) => {
      const userToRemove = getUserToRemove(connectedUsers, socketId);
      connectedUsers.delete(userToRemove);
    };
    const getUser = (userId) => {
      return connectedUsers.get(userId);
    };

    io.on("connection", (socket) => {
      const { clientId } = socket.handshake.query;
      if (clientId !== "undefined") {
        addUser(clientId, socket.id);
      }
      console.log("Connected users", connectedUsers);

      /**
       * Send a notification to consumer using the following template (JSON):
       {
        "id": "sdgdgfdgfd", //UUID string
        "toUser": 2,
        "type": "warning", //must be "warning", "error", "info", "success"
        "reason": "Attention",
        "description": "This is a warning message. Pay attention!"
      }
       */

      bindings.forEach(async ({ destination }) => {
        await channel.consume(destination, (notification) => {
          const toUserId = String(
            JSON.parse(notification.content.toString()).toUser
          );
          const targetUser = getUser(toUserId);
          if (targetUser) {
            const message = {
              ...notification,
              content: JSON.parse(notification.content.toString()),
            };
            io.to(targetUser).emit(
              BROKER_RECEIVED_NOTIFICATION,
              JSON.stringify(message)
            );
          }
        });
      });
      socket.on(CLIENT_RECEIVED_NOTIFICATION, (notification1) => {
        channel.ack(notification1);
      });

      socket.on("disconnect", () => {
        removeUser(socket.id);
        console.log("Connected users after leave", connectedUsers);
      });
    });
  } catch (e) {
    console.log(e);
  }
};
