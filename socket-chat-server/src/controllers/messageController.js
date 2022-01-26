import { Message } from "../models/models.js";

class MessageController {
  async add(req, res) {
    const { producerId, consumerId, conversationId, text } = req.body;
    console.log(req.body);

    const addedMessage = await Message.create({
      producerId,
      consumerId,
      conversationId,
      text,
    });

    return res.status(200).json(addedMessage);
  }

  async getByConversationId(req, res) {
    const { conversationId } = req.params;

    const conversationMessages = await Message.findAll({
      where: { conversationId },
    });

    return res.status(200).json(conversationMessages);
  }
}

export default new MessageController();
