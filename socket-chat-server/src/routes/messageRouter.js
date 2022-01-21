import { Router } from "express";
import MessageController from "../controllers/messageController.js";

const router = Router();

router.post("/add", MessageController.add);
router.get("/:conversationId", MessageController.getByConversationId);

export default router;
