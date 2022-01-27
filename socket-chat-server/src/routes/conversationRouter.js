import { Router } from "express";
import ConversationController from "../controllers/conversationController.js";

const router = Router();

router.post("/add", ConversationController.add);
router.get("/:userId", ConversationController.getByConversationsByUserId);
router.get(
  "/find/:producerId/:consumerId",
  ConversationController.getByConversationsOfTwoUsers
);

export default router;
