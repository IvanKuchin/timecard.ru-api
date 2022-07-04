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

// Timecard line structure
//
// swagger:model timecard_line
type timecard_line struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Number of reported hours per day, separated by comma
	// in: body
	//
	// required: true
	Row string
	// Task hours reported on
	// in: body
	//
	// required: true
	Tasks []task
}

// Timecard structure
//
// swagger:model timecard
type timecard struct {
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
	Approve_date string

	// SoW
	// in: body
	//
	// required: true
	Contract_sow []sow

	// SoW ID
	// in: body
	//
	// required: true
	Contract_sow_id string
	// Expected payment day (YYYY-MM-DD)
	// will be non-zero only if documents been delivered (see originals_received_date)
	// in: body
	//
	// required: true
	Expected_pay_day string
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
	// Reporting period start (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Period_start string
	// Reporting period end (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Period_end string
	// Timecard status (approved / saved / submit / rejected / pending_approve)
	// in: body
	//
	// required: true
	Status string
	// Time entries
	// in: body
	//
	// required: true
	Lines []timecard_line
	// Timestamp when expense was submitted to the system
	// If status is "saved" then it is 0
	// in: body
	//
	// required: true
	Submit_date string
}

// Returns list of SoW (Statement of Works)
type timecardServerResponse struct {
	// Result status
	// in: body
	Result string
	// Array of SoW
	// in: body
	Timecards []timecard
}

// Returns list of timecards
// swagger:response timecardServerResponseWrapper
type timecardServerResponseWrapper struct {
	// Result status
	// in: body
	Body timecardServerResponse
}
