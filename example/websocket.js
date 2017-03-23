/**
 * Created by aron on 17-3-23.
 */

window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");
    var roomIdTxt = document.getElementById("room_id")

    function appendLog(item) {
        var doScroll = log.scrollTop === log.scrollHeight - log.clientHeight;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("form").onsubmit = function () {
        var content = msg.value;
        content = JSON.stringify({content: content});
        var message = {protocol_id: 200, body: content};
        message = JSON.stringify(message);
        console.log(message);
        conn.send(message);
        msg.value = "";
        return false;
    };
    document.getElementById("enter_room").onclick = function () {
        var roomId = parseInt(roomIdTxt.value);
        var body = JSON.stringify({room_id: roomId});
        var message = {protocol_id: 100, body: body};
        message = JSON.stringify(message);
        console.log(message);
        conn.send(message);
        msg.value = "";
        return false;
    }
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://127.0.0.1:8080/websocket");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};