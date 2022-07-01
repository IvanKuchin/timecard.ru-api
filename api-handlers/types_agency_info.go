// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//  Schemes: http
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

// Holiday structure
//
// swagger:model holiday
type holiday struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reference to company id utilizing this holiday
	// in: body
	//
	// required: true
	Agency_company_id string
	// Holiday date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Date string
	// Holiday name
	// in: body
	//
	// required: true
	Title string
	// Full day off or half (full / half)
	// in: body
	//
	// required: true
	Type string
}

// Cost center structure
//
// swagger:model cost_center
type cost_center struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Refernece to agency
	// in: body
	//
	// required: true
	Agency_company_id string
	// Refernece to user assigned cost center to agency
	// in: body
	//
	// required: true
	Assignee_user_id string
	// Description
	// in: body
	//
	// required: true
	Description string
	// Contract expiration date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	End_date string
	// Contract number
	// in: body
	//
	// required: true
	Number string
	// Contract signing date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Sign_date string
	// Contract start date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Start_date string
	// Cost center company name
	// in: body
	//
	// required: true
	Title string
	// Company object
	// in: body
	//
	// required: true
	Companies []company
}

// Cost center assignment structure
//
// swagger:model cost_center_assignment
type cost_center_assignment struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Refernece to user assigned cost center to agency
	// in: body
	//
	// required: true
	Assignee_user_id string
	// Reference to cost center
	// in: body
	//
	// required: true
	Cost_center_id string
	// Reference to customer that will be assigned to cost senter
	// in: body
	//
	// required: true
	Timecard_customer_id string
}

// BT allowance structure
//
// swagger:model bt_allowance
type bt_allowance struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reference to company ID
	// in: body
	//
	// required: true
	Agency_company_id string
	// Amount in domestic currency must be paid to subcontractor before travel
	// in: body
	//
	// required: true
	Amount string
	// Destination country
	// in: body
	//
	// required: true
	Countries []country
}

// Flight destination structure
//
// swagger:model flight_destination
type flight_destination struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Country code
	// in: body
	//
	// required: true
	Abbrev string
	// Country name
	// in: body
	//
	// required: true
	Title string
}

// Airfare limit structure
//
// swagger:model airfare_limitation
type airfare_limitation struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Max amount per ticket
	// in: body
	//
	// required: true
	Limit string
	// Flight source airport
	// in: body
	//
	// required: true
	From []flight_destination
	// Flight destination airport
	// in: body
	//
	// required: true
	To []flight_destination
}

// Agency structure
//
// swagger:model agency
type agency struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string

	// Templates Business Trips build from
	// in: body
	//
	// required: true
	Bt_expense_templates []expense_template

	// Agency info
	// in: body
	//
	// required: true
	Companies []company

	// Limitations on ordering airfares
	// in: body
	//
	// required: true
	Airfare_limitations_by_direction []airfare_limitation

	// Allowance paid to subcontractor prior to travel
	// in: body
	//
	// required: true
	Allowances []bt_allowance

	// Cost center assignments
	// in: body
	//
	// required: true
	Cost_center_assignment []cost_center_assignment

	// Cost center list
	// in: body
	//
	// required: true
	Cost_centers []cost_center

	// Task list that subcontractors are working on
	// in: body
	//
	// required: true
	Tasks []task

	// National calendar
	// in: body
	//
	// required: true
	Holiday_calendar []holiday
}

// Returns list of BT (business trips)
type agency_infoServerResponse struct {
	// Result status
	// in: body
	Result string
	// Company info
	// in: body
	Agencies []agency
}

// Returns list of business trips
// swagger:response agency_infoServerResponseWrapper
type agency_infoServerResponseWrapper struct {
	// Result status
	// in: body
	Body agency_infoServerResponse
}
