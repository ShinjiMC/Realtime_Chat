const socket = new WebSocket("ws://localhost:3000/ws");

const connect = (cb: (msg: MessageEvent) => void) => {
  console.log("connecting");
  socket.onopen = () => {
    console.log("Successfully Connected");
  };
  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };
  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };
  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

const sendMsg = (msg: string | ArrayBufferLike | Blob | ArrayBufferView) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
