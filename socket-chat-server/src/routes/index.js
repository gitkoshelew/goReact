import { Router } from "express";
import messageRouter from "./messageRouter.js";
import conversationRouter from "./conversationRouter.js";

const router = Router();

router.use("/messages", messageRouter);
router.use("/conversations", conversationRouter);

export default router;
