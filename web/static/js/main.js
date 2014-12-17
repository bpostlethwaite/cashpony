var conn = new WebSocket('ws://localhost:8888/ws');
var Table = require('./table');
var tab = new Table(document.querySelector('table'), conn);
var Plotly = require('./plotly');
var plot = document.querySelector('#plot').contentWindow;
var plotly = new Plotly(plot);


window.tab = tab;

conn.onclose = function(evt) {
    console.log('closing connection');
};

conn.onmessage = function(evt) {
    var msg = JSON.parse(evt.data);
    console.log("message:", msg);
    tab.update(msg.record);
};


var data = {
    x: [[1,2,3,4,5]],
    y: [[15,15,26,16,8]],
    type: 'bar'
};


plotly.on('ready', function () {

    //tab.setReporter(plotly);

    plotly.plot(data);
});


window.plotly = plotly;
