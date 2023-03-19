package stdr

var ValidatorState = map[uint64]string{
	0: "Initialized",
	1: "Invalid Signature",
	2: "Front Run",
	3: "Pre Deposit",
	4: "Deposited",
	5: "In Activation Queue",
	6: "Active",
	7: "In Exit Queue",
	8: "Exited",
	9: "Withdrawn",
}
