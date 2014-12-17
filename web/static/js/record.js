var EventEmitter = require('events').EventEmitter;
var inherits = require('util').inherits;

function Record (data) {

    EventEmitter.call(this);

    var _elem = document.createElement('tr');
    var configured = false;

    _elem.innerHTML = this.template;
    _elem.id = this.id = data.id;

    this.name = '';
    this.date = '';
    this.debit = null;
    this.label = '';

    this._elem = _elem;

    // Add labels then inject!
    this.addLabels( this.labels );
    this.inject( data );

    var self = this;
    // Add dynamic interation
    this._elem.querySelector('select')
        .addEventListener('change', function (ev) {
            self.label = ev.target.value;
            self.emit('update', {label: self.label});
        });
}

inherits(Record, EventEmitter);

module.exports = Record;

Record.prototype.labels = [
    "Household",
    "Dining",
    "Transit",
    "Food",
    "Clothing",
    "Misc",
    "Unknown"
];

Record.prototype.inject = function (data) {
    this.name = data.name;
    this.date = data.date;
    this.debit = data.debit;
    this.label = data.label;

    var fields = this.getFields();

    this.injectDate (this.date, fields[0]);
    this.injectTrans(this.name, fields[1]);
    this.injectDebit(this.debit, fields[2]);
    this.injectLabel(this.label, fields[3]);
};

Record.prototype.getFields = function () {
    var row = this._elem;
    var tds = row.querySelectorAll('td');
    return [].slice.call(tds);
};

Record.prototype.injectDate = function (dateStr, td) {
    if (typeof dateStr !== 'string') return;
    var d = new Date(dateStr);
    if (isNaN(d.getTime())) return;
    var date = d.toLocaleDateString();
    if (td.textContent === date) return;

    td.textContent = date;
};

Record.prototype.injectTrans = function (txt, td) {
    if (typeof txt !== 'string') return;
    if (td.textContent === txt) return;
    td.textContent = txt;
};

Record.prototype.injectDebit = function (num, td) {
    if (typeof num !== 'number') return;
    var str = '$' + num;
    if (td.textContent === str) return;
    td.textContent = str;
};

Record.prototype.injectLabel = function (option, td) {
    if (typeof option !== 'string' ||
        this.labels.indexOf(option) === -1) {
        return;
    }

    var select = td.querySelector('select');
    select.value = option;
};

Record.prototype.clearLabels = function () {
    var select = this._elem.querySelector('select');
    for (var i = select.length-1; i >= 0; i--) select.remove(i);
};

Record.prototype.addLabels = function (options) {
    var select = this._elem.querySelector('select');

    options.forEach( function (option) {
        var opt = document.createElement("option");
        opt.value = option;
        opt.text = option;

        select.add(opt);
    });
};

Record.prototype.toString = function () {
    function formatter(key, value) {
        if (key[0] === '_') return undefined;
        else return value;
    }
    return JSON.stringify(this, formatter, 2);
};

Record.prototype.template =
    '<td></td>' +
    '<td></td>' +
    '<td></td>' +
    '<td>' +
    '  <select class="no-border-white">' +
    '  </select>' +
    '</td>';
