import { useState } from "react";
import "./App.css";

// Vite はトランスパイル時に import.meta.env のプロパティを VITE_ から始まる環境変数に置換する
// これを利用して本番環境と開発環境で Fetch API のリクエスト先を切り替えられる
// 参考: https://ja.vitejs.dev/guide/env-and-mode.html
const API_ENDPOINT = import.meta.env.VITE_API_ENDPOINT!;

type Message = {
  id: number;
  content: string;
};

function App() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputContent, setInputContent] = useState<string>("");

  async function updateMessages() {
    const response = await fetch(API_ENDPOINT + "/messages");
    const json = await response.json();
    const messages = json as Message[];
    setMessages(messages);
  }

  setInterval(updateMessages, 5 * 1000);
  window.onload = updateMessages;

  let i = 0;
  // FIXME: the browser starts to lag after like 1 min.
  // I'm not very familiar with JS, but isn't it GC'd?
  return (
    <>
      <main>
        <h1>TypeScript Frontend + Go Backend</h1>
        <ul>
          {messages.map((message: Message) => (
            <li key={i++}>{message.content}</li>
          ))}
        </ul>
        <input value={inputContent} placeholder="メッセージを入力" onChange={(e) => {
          setInputContent(e.target.value);
        }}/>
        <button
          type="submit"
          disabled={inputContent === ""}
          onClick={async () => {
            const copy = inputContent;
            setInputContent("");
            await sendMessage(copy);
            await updateMessages();
          }}
        >
          送信
        </button>
      </main>
    </>
  );
}

async function sendMessage(content: string) {
  await fetch(API_ENDPOINT + "/send", {
    method: "post",
    body: JSON.stringify({
      content,
    }),
  });
}

export default App;
