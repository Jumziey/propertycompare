package property

//Loan describes a loan with rent, amount and amortization
type Loan struct {
	//Amount is the size of the loan
	Amount float64
	//Rent percent on the part of the down payment you borrow, yearly
	Rent float64
	//Amortization percent on the part of the down payment you borrow, yearly
	Amortization float64
}

//Mortgage represent a property mortgage.
type Mortgage Loan
