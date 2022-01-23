import { Router } from "express";
import ConversationController from "../controllers/conversationController.js";

const router = Router();

router.post("/add", ConversationController.add);
router.get("/:userId", ConversationController.getByConversationsByUserId);
router.get(
  "/find/:firstUser/:secondUser",
  ConversationController.getByConversationsOfTwoUsers
);

export default router;
