var EventEmitter = require('events').EventEmitter;
var inherits = require('util').inherits;

function Plotly(plot) {

    EventEmitter.call(this);

    this.plotly = plot;

    this.init();
}

inherits(Plotly, EventEmitter);

module.exports = Plotly;

Plotly.prototype.init = function () {

    window.removeEventListener('message');

    var self = this;
    var plotly = this.plotly;
    var pinger = setInterval(function(){
        console.log('posting ping');
        plotly.postMessage({ping: true}, 'https://plot.ly');
    }, 1000);



    window.addEventListener('message', function(e) {
        var message = e.data;
        if(message==='pong') {
            clearInterval(pinger);
            self.emit('ready');
        }
    });
};

Plotly.prototype.plot = function (data) {

    this.plotly.postMessage({restyle: data}, 'https://plot.ly');

};
