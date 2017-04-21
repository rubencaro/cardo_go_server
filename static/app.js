var source = new EventSource('/events');
source.onopen = function (event) {
  console.log("eventsource connection open");
};
source.onerror = function (event) {
  if (event.target.readyState === 0) {
    console.log("reconnecting to eventsource");
  } else {
    console.log("eventsource error");
  }
};
source.onmessage = function (event) {
  var txt = document.createTextNode(event.data);
  var div = document.createElement("div");
  div.appendChild(txt);

  var first = document.getElementById("messages").firstChild;
  document.getElementById("messages").insertBefore(div, first);
};

var add = function () {
  var body = { "msg": document.getElementById("text").value };
  fetch("/add", {
    "method": "POST",
    "headers": { "content-type": "application/json" },
    "body": JSON.stringify(body)
  });
}