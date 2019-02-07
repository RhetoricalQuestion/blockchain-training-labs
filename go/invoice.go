package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

//"github.com/hyperledger/fabric/core/chaincode/lib/cid"

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the Invoice structure
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

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "raiseInvoice" {
		return s.raiseInvoice(APIstub, args)
	} else if function == "goodsReceived" {
		return s.goodsReceived(APIstub)
	} else if function == "bankPaymentToSupplier" {
		return s.bankPaymentToSupplier(APIstub, args)
	} else if function == "oemRepaymentToBank" {
		return s.oemRepaymentToBank(APIstub, args)
	} else if function == "displayAllInvoices" {
		return s.displayAllInvoices(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}
