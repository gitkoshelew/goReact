import { Sequelize } from "sequelize";
import dotenv from "dotenv";

dotenv.config();

export default new Sequelize("goreact", "user", "userpass", {
  dialect: "postgres",
  host: "localhost",
  port: 8081,
});
