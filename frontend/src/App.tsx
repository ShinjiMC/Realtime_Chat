import { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import ChatHistory from "./components/chatHistory";
import Header from "./components/header";

function App() {
  const [chatHistory, setChatHistory] = useState<{ data: string }[]>([]);

  const send = () => {
    console.log("hello");
    sendMsg("hello");
  };

  useEffect(() => {
    const handleMessage = (msg: { data: string }) => {
      console.log("New Message");
      setChatHistory((prevChatHistory) => [...prevChatHistory, msg]);
    };
    connect(handleMessage);
  }, []);

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <button onClick={send}>Hit</button>
    </div>
  );
}

export default App;
