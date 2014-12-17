function Message (spec) {

    if (!spec) spec = {};

    this.record = spec.record;
    this.labelupdate = (spec.labelupdate === undefined) ? false : spec.labelupdate;
    this.msg = spec.msg || '';
}

module.exports = Message;


Message.prototype.serialize = function () {

    function formatter (key, value) {
        if (key[0] === '_') return undefined;
        else return value;
    }

    return JSON.stringify(this, formatter, 2);
};
