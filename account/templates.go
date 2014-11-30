package account

type csvTemplate struct {
	DateField        int
	DateFormat       string
	TransactionField int
	DebitField       int
	CreditField      int
}

var TDCC = csvTemplate{
	DateField:        0,
	DateFormat:       "01/02/2006",
	TransactionField: 1,
	DebitField:       2,
	CreditField:      3,
}
