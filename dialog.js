var socket;
var smsDiv;

function Init(sender, reciever){
 socket = new WebSocket("ws://localhost:8080/peer/" + sender + "/" + reciever);

 socket.onopen = function(event) {
    console.log("Соединение установлено");
  };

  socket.onmessage = function(event) {
    var message = JSON.parse(event.data)
    console.log("Получено сообщение: " + message);
    const newTr = document.createElement("tr");
    const usernameTd = document.createElement("td");
    const messageTd = document.createElement("td");
    usernameTd.innerHTML = message.username;
    messageTd.innerHTML = message.message;
    // newTr.innerHTML = event.data.toString();
    newTr.appendChild(usernameTd)
    newTr.appendChild(messageTd)
    smsDiv.appendChild(newTr);
  };

  socket.onclose = function(event) {
    console.log("Соединение закрыто");
  };
}

function sendMessage() {
  var message = document.getElementById("message").value;
  socket.send(message);
}

document.addEventListener("DOMContentLoaded", function() {
  smsDiv = document.getElementById("sms");
  console.log(document);
  console.log(smsDiv);
});


