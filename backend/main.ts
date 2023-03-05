import express from "express";
import cors from "cors";
import { PrismaClient } from "@prisma/client";

const client = new PrismaClient();
const app = express();

// 別ドメインから Fetch API を用いてリクエストを送信可能にするために必要
// WEB_ORIGIN に設定したドメインからのリクエストのみ許可する
// 参考: https://developer.mozilla.org/ja/docs/Web/HTTP/CORS
app.use(cors({ origin: process.env.WEB_ORIGIN }));

app.use(express.json());

app.get("/messages", async (request, response) => {
  response.json(await client.message.findMany());
});

app.post("/send", async (request, response) => {
  await client.message.create({ data: { content: request.body.content } });
  response.send();
});

app.listen(3000);
