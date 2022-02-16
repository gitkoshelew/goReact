import amqp from "amqplib";
import rabbitDefinitions from "../../../definitions.json" assert { type: "json" };
import sequelize from "../../../db.js";
import { QueryTypes } from "sequelize";
import moment from "moment";
import { v1 } from "uuid";

const rabbitSettings = {
  protocol: "amqp",
  hostname: process.env.RABBIT_HOST,
  port: +process.env.RABBIT_PORT,
  username: process.env.RABBIT_USER,
  password: process.env.RABBIT_PASSWORD,
  vhost: "/",
  authMechanism: ["PLAIN", "AMQPLAIN", "EXTERNAL"],
};

const NOTIFICATIONS_EXCHANGE = rabbitDefinitions.exchanges.find(
  (exchange) => exchange.name === "NOTIFICATIONS_EXCHANGE"
).name;

const hoursSearchInterval = 6;

export const startNotificationProducer = async () => {
  try {
    const connection = await amqp.connect(rabbitSettings);
    const channel = await connection.createChannel();

    await channel.assertExchange(NOTIFICATIONS_EXCHANGE, "direct", {
      durable: true,
    });

    const tomorrow = moment().add(1, "days").format("YYYY-MM-DD");
    setInterval(async () => {
      const usersForNotification = await sequelize.query(
        `SELECT booking.pet_id, pet.user_id FROM booking INNER JOIN pet ON booking.pet_id = pet.id WHERE booking.start_date = '${tomorrow}'`,
        {
          type: QueryTypes.SELECT,
        }
      );

      usersForNotification.forEach((user) => {
        const notification = {
          id: v1(),
          toUser: user.user_id,
          type: "warning",
          reason: "You have book for tomorrow",
          description:
            "Dear customer! You have booked a room in our hotel tomorrow",
        };
        channel.publish(
          NOTIFICATIONS_EXCHANGE,
          "notification_warning",
          Buffer.from(JSON.stringify(notification))
        );
      });
    }, 7000);
  } catch (e) {
    console.log(e);
  }
};
