// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//  Schemes: http, https
//  BasePath: /
//  Version: 1.0.7
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package apihandlers

// Approval structure
//
// swagger:model approval
type approval struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Approver ID
	// in: body
	//
	// required: true
	Approver_id string
	// Business trip ID
	// in: body
	//
	// required: true
	Bt_id string
	// Comment (used if decision is reject)
	// in: body
	//
	// required: true
	Comment string
	// Decision (approve / reject)
	// in: body
	//
	// required: true
	Decision string
}

// BT expense line structure
//
// swagger:model expense_line
type expense_line struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reference to expense
	// in: body
	//
	// required: true
	Bt_expense_id string
	// Reference to template
	// in: body
	//
	// required: true
	bt_expense_line_template_id string
	// Comment
	// in: body
	//
	// required: true
	Comment string
	// Single expense value:
	//   if type is input, then value is simple text
	//   if type is pdf, then value is URL to pdf-file
	//   if type is image, then value is URL to image-file
	// in: body
	//
	// required: true
	Value string
}

// BT expense structure
//
// swagger:model expense
type expense struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Task hours reported on
	// in: body
	//
	// required: true
	Bt_expense_line []expense_line
	// Task hours reported on
	// in: body
	//
	// required: true
	Bt_expense_templates []expense_template
	// Reference to BT
	// in: body
	//
	// required: true
	Bt_id string
	// Comment
	// in: body
	//
	// required: true
	Comment string
	// Currency, if paid not RUR
	// in: body
	//
	// required: true
	Currency_name string
	// Currency nominal is a basis to calculate rate exchange
	// in: body
	//
	// required: true
	Currency_nominal string
	// Currency value is a basis to calculate rate exchange
	// in: body
	//
	// required: true
	Currency_value string
	// Expense date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Date string
	// Flag will be set if rate exchange pointed by the subcontractor is the same as Central Bank rate exchange at a date of expense
	// in: body
	//
	// required: true
	Is_cb_currency_rate string
	// Payment type (cash / credit card)
	// in: body
	//
	// required: true
	Payment_type string
	// Transaction in domestic currency
	// in: body
	//
	// required: true
	Price_domestic string
	// Transaction in foreign currency
	// in: body
	//
	// required: true
	Price_foreign string
}

// BT (business trip) structure
//
// swagger:model bt
type bt struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string

	// Timestamp when timecard been approved
	// in: body
	//
	// required: true
	Approvals []approval
	// Array of approvers
	// in: body
	//
	// required: true
	Approvers []approver
	// Timestamp when timecard been approved
	// in: body
	//
	// required: true
	Approve_date string

	// SoW
	// in: body
	//
	// required: true
	Sow []sow

	// SoW ID
	// in: body
	//
	// required: true
	Contract_sow_id string
	// Cutomer to visit
	// in: body
	//
	// required: true
	Customers []customer
	// Start date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Date_start string
	// End date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Date_end string
	// Expected payment day (YYYY-MM-DD)
	// will be non-zero only if documents been delivered (see originals_received_date)
	// in: body
	//
	// required: true
	Expected_pay_day string
	// Expenses to reimburse
	// in: body
	//
	// required: true
	Expenses []expense
	// URL to doenload invoice package
	// in: body
	//
	// required: true
	Invoice_filename string
	// Day when paper copy of documents been received (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Originals_received_date string
	// Payment day (YYYY-MM-DD)
	// will be non-zero only if payment been marked as happened
	// in: body
	//
	// required: true
	Payed_date string
	// BT destination
	// in: body
	//
	// required: true
	Place string
	// BT purpose
	// in: body
	//
	// required: true
	Purpose string
	// BT status (approved / saved / submit / rejected / pending_approve)
	// in: body
	//
	// required: true
	Status string
	// Timestamp when expense was submitted to the system
	// If status is "saved" then it is 0
	// in: body
	//
	// required: true
	Submit_date string
}

// Returns list of BT (business trips)
type btServerResponse struct {
	// Result status
	// in: body
	Result string
	// Array of SoW
	// in: body
	Bt []bt
}

// Returns list of business trips
// swagger:response btServerResponseWrapper
type btServerResponseWrapper struct {
	// Result status
	// in: body
	Body btServerResponse
}
