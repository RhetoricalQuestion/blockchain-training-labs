package main

//"github.com/hyperledger/fabric/core/chaincode/lib/cid"

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Invoice struct {
	InvoiceNumber   int    `json:"invoiceNumber"`
	BilledTo        string `json:"billedTo"`
	InvoiceDate     string `json:"invoiceDate"`
	InvoiceAmount   int    `json:"invoiceAmount"`
	ItemDescription string `json:"itemDescription"`
	GR              bool   `json:"gr"`     //[Y/N]
	IsPaid          bool   `json:"isPaid"` //[Y/N]
	PaidAmount      int    `json:"paidAmount"`
	Repaid          bool   `json:"repaid"` //[Y/N]
	RepaymentAmount int    `json:"repaymentAmount"`
}
