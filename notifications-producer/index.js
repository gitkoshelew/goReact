import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import sequelize from "./db.js";
import { startNotificationProducer } from "./src/services/notifications/producer.js";

dotenv.config();

const app = express();

app.use(
  cors({
    origin: ["http://localhost:3000", "http://localhost:3001"],
    credential: true,
  })
);
app.use(express.json());

const PORT = process.env.PORT || 5005;

app.get("/ping", (req, res) => {
  res.send("Working!");
});

const start = async () => {
  try {
    await sequelize.authenticate();
    await sequelize.sync();

    await startNotificationProducer();

    app.listen(PORT, () => console.log(`Server starts on port ${PORT}`));
  } catch (e) {
    console.log(e);
  }
};

start();
