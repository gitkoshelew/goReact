import { Conversation } from "../models/models.js";
import { Op } from "sequelize";

class ConversationController {
  async add(req, res) {
    const { senderId, receiverId } = req.body;

    const addedConversation = await Conversation.create({
      members: [senderId, receiverId],
    });

    return res.status(200).json(addedConversation);
  }

  async getByConversationsByUserId(req, res) {
    const { userId } = req.params;

    const userConversations = await Conversation.findAll({
      where: { members: { [Op.contains]: [userId] } },
    });

    return res.status(200).json(userConversations);
  }

  async getByConversationsOfTwoUsers(req, res) {
    const { firstUser, secondUser } = req.params;

    const usersConversations = await Conversation.findOne({
      where: { members: { [Op.contains]: [firstUser, secondUser] } },
    });

    return res.status(200).json(usersConversations);
  }
}

export default new ConversationController();
