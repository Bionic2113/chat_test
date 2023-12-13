var xhr = new XMLHttpRequest();

function Add(event, goods_id) {
  event.preventDefault();
  // var goods_id = event.target.getAttribute('data-id');
  // var user = JSON.stringify({
  //   surname: document.getElementById('surname').value,
  //   name: document.getElementById('name').value,
  //   note: document.getElementById('note').value.toString(),
  // });
  xhr.open("POST", "/edit?goods_id="+goods_id+"&action=add");
  xhr.setRequestHeader("Content-type", "application/text; charset=utf-8");
  console.log("Запрос сейчас будет отправлен. Goods_id: " + goods_id)
  xhr.send();
  console.log("Запрос отправлен")
  xhr.onreadystatechange = check;
}
function Reduce(event, goods_id) {
  event.preventDefault();
  // var goods_id = event.target.getAttribute('data-id');
  xhr.open("POST", "/edit?goods_id="+goods_id+"&action=reduce");
  xhr.setRequestHeader("Content-type", "application/text; charset=utf-8");
  console.log("Запрос сейчас будет отправлен")
  xhr.send();
  console.log("Запрос отправлен")
  xhr.onreadystatechange = check;
}


function check() {
  if (xhr.readyState == 4 && xhr.status == 200) {
    // document.body.innerHTML = xhr.responseText

    // этот вариант добавлял хтмл в заместо таблицы.
    // получался хтмл в хтмл)
    // var table = document.getElementById('table')
    // table.innerHTML = xhr.responseText //responseXML.getElementById('table')
    
    // document.getElementsByClassName('edit')[0].style = 'display: none'
    // document.getElementsByClassName('view')[0].style = 'display: block'
  
    console.log("дейтсвие выполнено")
  }
}

