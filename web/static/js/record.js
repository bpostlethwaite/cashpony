function Record (id) {

    var row = document.createElement('tr');

    row.id = this.id = id;
    row.innerHTML = this.template;

    this.row = row;
}

module.exports = Record;


Record.prototype.inject = function (data) {
    var date   = data.Date;
    var trans  = data.Transaction;
    var debit  = data.Debit;
    var label = data.Label;

    var fields = this.getFields();

    this.injectDate(date, fields[0]);
    this.injectTrans(trans, fields[1]);
    this.injectDebit(debit, fields[2]);
    this.injectLabel(fields, fields[3]);
};

Record.prototype.injectDate = function (dateStr, td) {
    if (typeof dateStr !== 'string') return;
    var d = new Date(dateStr);
    if (isNaN(d.getTime())) return;
    var date = d.toLocaleString();
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
    if (typeof option === 'string') return;

    var select = this.row.querySelector('select');
    if (option in this.selectOptions) {
        select.value = option;
    }
};

Record.prototype.getFields = function () {
    var row = this.row;
    var tds = row.querySelectorAll('td');
    return [].slice.call(tds);
};

Record.prototype.clearSelection = function () {
    var select = this.row.querySelector('select');
    for (var i = select.length-1; i >= 0; i--) select.remove(i);
};

Record.prototype.addOptions = function (options) {
    var select = this.row.querySelector('select');

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
