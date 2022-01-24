import sequelize from "../../db.js";
import { DataTypes } from "sequelize";

export const Message = sequelize.define("message", {
  id: {
    type: DataTypes.UUID,
    primaryKey: true,
    defaultValue: DataTypes.UUIDV1,
  },
  senderId: { type: DataTypes.INTEGER, allowNull: false },
  receiverId: { type: DataTypes.INTEGER, allowNull: false },
  conversationId: { type: DataTypes.INTEGER, allowNull: false },
  text: { type: DataTypes.STRING, allowNull: false },
});

export const Conversation = sequelize.define("conversation", {
  id: {
    type: DataTypes.INTEGER,
    primaryKey: true,
    autoIncrement: true,
  },
  members: { type: DataTypes.ARRAY(DataTypes.INTEGER) },
});
