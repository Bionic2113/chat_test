let socket: WebSocket;
let smsDiv: HTMLElement;

function Init(sender: string, reciever: string){
 socket = new WebSocket("ws://localhost:8080/peer/" + sender + "/" + reciever);
}

function sendMessage() {
  const message = (document.getElementById("message") as HTMLInputElement).value;
  socket.send(message);
}

document.addEventListener("DOMContentLoaded", function() {
  smsDiv = document.getElementById("sms") as HTMLElement;
  console.log(document);
  console.log(smsDiv);

  socket.onopen = function(event) {
    console.log("Соединение установлено");
  };

  socket.onmessage = function(event) {
    console.log("Получено сообщение: " + event.data);
    const newSpan = document.createElement("p");
    newSpan.innerHTML = event.data.toString();
    smsDiv.appendChild(newSpan);
  };

  socket.onclose = function(event) {
    console.log("Соединение закрыто");
  };
});
