import { Sequelize } from "sequelize";
import dotenv from "dotenv";
dotenv.config();

export default new Sequelize(
  process.env.CHAT_DB_NAME,
  process.env.CHAT_DB_USER,
  process.env.CHAT_DB_PASSWORD,
  {
    dialect: "postgres",
    host: process.env.CHAT_DB_HOST,
    port: process.env.CHAT_DB_PORT,
  }
);
