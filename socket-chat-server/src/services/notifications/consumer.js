import amqp from "amqplib";

const rabbitSettings = {
  protocol: "amqp",
  hostname: "rabbitmq",
  port: 5672,
  username: "guest",
  password: "guest",
  vhost: "/",
  authMechanism: ["PLAIN", "AMQPLAIN", "EXTERNAL"],
};

const NOTIFICATIONS_QUEUE = "NOTIFICATIONS_QUEUE";
const BROKER_RECEIVED_NOTIFICATION = "BROKER_RECEIVED_NOTIFICATION";

export const startNotificationConsumer = async (io) => {
  try {
    // Не коннектится к шине, после этой строчки сразу падает в catch
    const connection = await amqp.connect(rabbitSettings);
    console.log("Connection_____", connection);

    const channel = await connection.createChannel();
    console.log("Channel_____", channel);

    await channel.assertQueue(NOTIFICATIONS_QUEUE);
    /**
     * Send a notification to consumer using the following template (JSON):
     *{
      "toUser": 2
      "type": "warning", //must be "warning", "error", "info", "success"
      "reason": "Attention",
      "description": "This is a warning message. Pay attention!"
    }
     */
    await channel.consume(NOTIFICATIONS_QUEUE, (notification) => {
      io.emit(BROKER_RECEIVED_NOTIFICATION, notification.content.toString());
    });
  } catch (e) {
    console.log("Error________", e);
  }
};
