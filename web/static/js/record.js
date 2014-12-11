function Record (data) {

    var elem = document.createElement('tr');
    var configured = false;

    elem.innerHTML = this.template;
    elem.id = this.id = data.id;

    this.elem = elem;
    this.inject( data );
    this.addLabels( this.labels );
}

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
    var date   = data.date;
    var trans  = data.transaction;
    var debit  = data.debit;
    var label = data.label;

    var fields = this.getFields();

    this.injectDate(date, fields[0]);
    this.injectTrans(trans, fields[1]);
    this.injectDebit(debit, fields[2]);
    this.injectLabel(label, fields[3]);
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

    var select = this.elem.querySelector('select');
    select.value = option;
};

Record.prototype.getFields = function () {
    var row = this.elem;
    var tds = row.querySelectorAll('td');
    return [].slice.call(tds);
};

Record.prototype.clearLabels = function () {
    var select = this.elem.querySelector('select');
    for (var i = select.length-1; i >= 0; i--) select.remove(i);
};

Record.prototype.addLabels = function (options) {
    var select = this.elem.querySelector('select');

    options.forEach( function (option) {
        var opt = document.createElement("option");
        opt.value = option;
        opt.text = option;

        select.add(opt);
    });
};

Record.prototype.template =
    '<td></td>' +
    '<td></td>' +
    '<td></td>' +
    '<td>' +
    '  <select class="no-border-white">' +
    '  </select>' +
    '</td>';
