var socket = new WebSocket("ws://localhost:8080/ws");
// window.socketConnection = socket;
// Получаем элемент <div> по его id
// const smsDiv = document.getElementById("sms");

socket.onopen = function(event) {
  console.log("Соединение установлено");
  //socket.close()
};

// socket.onmessage = function(event) {
//   //console.log("Получено сообщение: " + event.data);
//   // Создаем новый элемент <span>
//   const newSpan = document.createElement("span");
//
//   // Устанавливаем текст для нового элемента
//   // newSpan.innerHTML = event.data.toString();
//
//   // // Добавляем новый элемент <span> в элемент <div>
//   // smsDiv.appendChild(newSpan);
// };

socket.onclose = function(event) {
  //socket.send("close hihi")
  console.log("Соединение закрыто");
};

// function sendMessage() {
//   var message = document.getElementById("message").value;
//   socket.send(message);
//   // socket.close()
// }
//
