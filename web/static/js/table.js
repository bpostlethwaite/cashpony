var Record = require('./record');
var Message = require('./message');

function Table (root, conn) {

    this.root = root;
    this.body = root.querySelector('tbody');
    this.conn = conn;
    this.records = {};
}

module.exports = Table;

Table.prototype.update = function (data) {

    if (!data || !data.id) return;

    var tbody = this.body;
    var rec = this.records[data.id];
    var conn = this.conn;

    if (rec) {
        rec.inject(data);
        return;
    }

    // else make a new record
    rec = new Record(data);

    // add to hash
    this.records[rec.id] = rec;

    // append the rec
    tbody.appendChild(rec._elem);

    // put this in Rec then just pass in a cb
    rec.on('update', function (update) {
        var msg = new Message({
            record: rec,
            labelupdate: 'label' in update
        });
        conn.send(msg.serialize());
    });
};
