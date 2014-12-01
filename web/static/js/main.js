var Record = require('./record');


var conn;
var table = document.querySelector('table');
var tbody = table.querySelector('tbody');

function msgRouter(msg) {
    var row = new Record(Math.random() + 'a');
    tbody.appendChild(row.row);

    row.inject(msg);
    console.log(msg);
}

if (window.WebSocket) {
    conn = new WebSocket('ws://localhost:8080/ws');
    conn.onclose = function(evt) {
        console.log('closing connection');
    };
    conn.onmessage = function(evt) {
        var msg = JSON.parse(evt.data);
        msgRouter(msg);
    };
} else {
    console.log(' no websockets ');
}


setTimeout( function () {
    conn.send(JSON.stringify({msg: 'HUZZAH!'}));
}, 1000);
