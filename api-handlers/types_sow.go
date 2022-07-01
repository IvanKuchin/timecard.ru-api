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

// Company structure
//
// swagger:model bank
type bank struct {
	// bank ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Account #
	// in: body
	//
	// required: true
	Account string
	// Address
	// in: body
	//
	// required: true
	Address string
	// Bank ID Code
	// in: body
	//
	// required: true
	Bik string
	// Bank location
	// in: body
	//
	// required: true
	Geo_zip_id []geo_zip
	// OKATO
	// in: body
	//
	// required: true
	Okato string
	// OKPO
	// in: body
	//
	// required: true
	Okpo string
	// Phone number
	// in: body
	//
	// required: true
	Phone string
	// Bank name
	// in: body
	//
	// required: true
	Title string
	// Avarage duration of completing transaction
	// in: body
	//
	// required: true
	Srok string
}

// Object describing country
//
// swagger:model country
type country struct {
	// country ID
	// same as country phone code
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Country
	// in: body
	//
	// required: true
	Title string
}

// Object describing geo-region
//
// swagger:model region
type region struct {
	// region ID, same as region id inside country
	// for instance: Moscow is 77
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Region
	// in: body
	//
	// required: true
	Title string
	// Country object
	// in: body
	//
	// required: true
	Country country
}

// Object describing city/town
//
// swagger:model locality
type locality struct {
	// Locality ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// City/town
	// in: body
	//
	// required: true
	Title string
	// Region object
	// in: body
	//
	// required: true
	Region region
}

// Object describing location by zip-code
//
// swagger:model geo_zip
type geo_zip struct {
	// Zip code ID (just an id in DB, not an actual zip-code)
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// zip code
	// in: body
	//
	// required: true
	Zip string
	// Locality object
	// in: body
	//
	// required: true
	Locality locality
}

// Company structure
//
// swagger:model company
type company struct {
	// company ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Company name
	// in: body
	//
	// required: true
	Name string
	// Company type
	// subcontractor or agency
	// in: body
	//
	// required: true
	Type string
	// Description
	// in: body
	//
	// required: true
	Description string
	// Bank account
	// in: body
	//
	// required: true
	Account string
	// Bank KPP
	// in: body
	//
	// required: true
	Kpp string
	// Bank OGRN
	// in: body
	//
	// required: true
	Ogrn string
	// TIN (Tax ID Number)
	// in: body
	//
	// required: true
	Tin string
	// If company works with VAT, then "Y", otherwise "N"
	// in: body
	//
	// required: true
	Vat string
	// Act prefix (full Act ID contains: Prefix + # + Postfix)
	// Prefix must be used for persistant parts of act
	// for instance: abc123def
	// ABC is a prefix
	// def is a postfix
	// in: body
	//
	// required: true
	Act_number_prefix string
	// Act #
	// Number automatically increases after each report period
	// in: body
	//
	// required: true
	Act_number string
	// Act postfix (full Act ID contains: Prefix + # + Postfix)
	// Prefix must be used for persistant parts of act
	// for instance: abc123def
	// abc is a prefix
	// def is a postfix
	// in: body
	//
	// required: true
	Act_number_postfix string
	// Bank credentials
	// in: body
	//
	// required: true
	Banks []bank
	// Ownership flag
	// 0 - means requestor doesn't own this company
	// 1 - means requestor owns this company
	// in: body
	//
	// required: true
	IsMine string
	// Legal address
	// in: body
	//
	// required: true
	Legal_address string
	// Object containing info about ZIP-code of legal company address
	// in: body
	//
	// required: true
	Legal_geo_zip []geo_zip
	// Mailing address
	// in: body
	//
	// required: true
	Mailing_address string
	// Object containing info about ZIP-code of mailing company address
	// in: body
	//
	// required: true
	Mailing_geo_zip []geo_zip
	// URL to company web-site
	// in: body
	//
	// required: true
	WebSite string
	// Logo folder
	// in: body
	//
	// required: true
	Logo_folder string
	// Logo filename
	// in: body
	//
	// required: true
	Logo_filename string
}

// Document that must be submitted to prove expense relevancy
//
// swagger:model expense_template
type expense_template struct {
	// template ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reference to template ID
	// in: body
	//
	// required: true
	Bt_expense_template_id string
	// Doc description
	// in: body
	//
	// required: true
	Description string
	// Type of document (pdf, image)
	// in: body
	//
	// required: true
	Dom_type string
	// Payment type (cash, credit card)
	// in: body
	//
	// required: true
	Payment string
	// Is this document mandatory for reimbursement Y/N
	// in: body
	//
	// required: true
	Required string
	// Document name
	// in: body
	//
	// required: true
	Title string
	// Description that will be displayed "on hover" event
	// in: body
	//
	// required: true
	Tooltip string
}

// Business trip template structure
//
// swagger:model bt_expense_template
type bt_expense_template struct {
	// template ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Reasoning of this expense
	// in: body
	//
	// required: true
	Agency_comment string
	// If Y, then this expense will be assigned to expense report
	// this is convenience field for operator to build expenses from template
	// in: body
	//
	// required: true
	Is_default string
	// If Y, then taxes will be added on top of expense cost
	// in: body
	//
	// required: true
	Taxable string
	// Expense
	// in: body
	//
	// required: true
	Title string
	// Documents that must be submitted to prove expense relevancy
	// in: body
	//
	// required: true
	Line_templates []expense_template
}

// Company occupation
//
// swagger:model company_position
type company_position struct {
	// template ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Occupancy
	// in: body
	//
	// required: true
	Title string
}

// User structure
//
// swagger:model user
type user struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Avatar URL
	// in: body
	//
	// required: true
	Avatar string
	// Birthday (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Birthday string
	// Country of citizenship (RU)
	// in: body
	//
	// required: true
	Citizenship_code string
	// Email
	// in: body
	//
	// required: true
	Email string
	// Travel passport exp date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Foreign_passport_expiration_date string
	// Travel passport number
	// in: body
	//
	// required: true
	Foreign_passport_number string
	// Flag of myself (yes if user object equal to logged in user)
	// in: body
	//
	// required: true
	IsMe string
	// First name (written in english)
	// this field used to book travels
	// in: body
	//
	// required: true
	First_name_en string
	// Middle name (written in english)
	// this field used to book travels
	// in: body
	//
	// required: true
	Middle_name_en string
	// Last name (written in english)
	// this field used to book travels
	// in: body
	//
	// required: true
	Last_name_en string
	// First name
	// in: body
	//
	// required: true
	Name string
	// Middle name
	// in: body
	//
	// required: true
	NameMiddle string
	// Last name
	// in: body
	//
	// required: true
	NameLast string
	// Authorithy issued native passport
	// in: body
	//
	// required: true
	Passport_issue_authority string
	// Date of issuing native passport (YYYY-MM-DD)
	// in: body
	//
	// required: true
	Passport_issue_date string
	// Number of native passport
	// in: body
	//
	// required: true
	Passport_number string
	// Series of native passport
	// in: body
	//
	// required: true
	Passport_series string
	// Is any BT approvals pending
	// in: body
	//
	// required: true
	Pending_approval_notification_bt string
	// Is any Service approvals pending
	// in: body
	//
	// required: true
	Pending_approval_notification_timecard string
	// Phone
	// in: body
	//
	// required: true
	Phone string
	// User sex
	// in: body
	//
	// required: true
	UserSex string
	// User type (agency or subcontractor)
	// in: body
	//
	// required: true
	UserType string
}

// Approver structure
//
// swagger:model approver
type approver struct {
	// ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Order in approval chain
	// in: body
	//
	// required: true
	Approver_order string
	// Auto-approve Y/N
	// in: body
	//
	// required: true
	Auto_approve string
	// Contract PSoW
	// in: body
	//
	// required: true
	Contract_psow_id string
	// Rate to calculate (subcontractor or agency)
	// in: body
	//
	// required: true
	Rate string
	// Approval entity (service or bt (business trip))
	// in: body
	//
	// required: true
	Type string
	// User object
	// in: body
	//
	// required: true
	Users []user
}

// PSoW structure
//
// swagger:model psow
type psow struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Business trip markup
	// in: body
	//
	// required: true
	Bt_markup string
	// Business trip markup type
	//   percent - percent from the whole trip
	//   fix - fixed amount
	// in: body
	//
	// required: true
	Bt_markup_type string
	// Occupation in the contract
	// in: body
	//
	// required: true
	Company_positions []company_position
	// Reference to SoW bound to current PSoW
	// in: body
	//
	// required: true
	Contract_sow_id string
	// Day rate toward cost center (usually subcontractors day rate + markup)
	// in: body
	//
	// required: true
	Day_rate string
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
	// Contract expiraton date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	End_date string
	// Contract number
	// in: body
	//
	// required: true
	Number string
	// Amount of working hours per day (extra hours will be highlihted as extrahours)
	// in: body
	//
	// required: true
	Working_hours_per_day string
}

// Task structure
//
// swagger:model task
type task struct {
	// Task ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Title
	// in: body
	//
	// required: true
	// min: 1
	Title string
	// List of projects this task assigned to
	// in: body
	//
	// required: true
	// min: 1
	Projects []project
}

// Project structure
//
// swagger:model project
type project struct {
	// Task ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Title
	// in: body
	//
	// required: true
	// min: 1
	Title string
	// List of customers this project assigned to
	// in: body
	//
	// required: true
	// min: 1
	Customers []customer
}

// Customer structure
//
// swagger:model customer
type customer struct {
	// Task ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Title
	// in: body
	//
	// required: true
	// min: 1
	Title string
}

// SoW structure
//
// swagger:model sow
type sow struct {
	// SoW ID
	// in: body
	//
	// required: true
	// min: 1
	Id string
	// Agency holding the contract
	// in: body
	//
	// required: true
	Agency_company_id []company
	// URL to download an agreement
	// in: body
	//
	// required: true
	Agreement_filename string
	// Array of persons to approve busines trips
	// in: body
	//
	// required: true
	Bt_approvers []approver
	// Array of templates used for business trip
	// in: body
	//
	// required: true
	Bt_expense_templates []bt_expense_template
	// Subcontractor occupation
	// in: body
	//
	// required: true
	Company_positions []company_position
	// Cost center that will be charged for subcontractor service
	// in: body
	//
	// required: true
	Cost_centers []company
	// Dayrate with subcontractor
	// in: body
	//
	// required: true
	Day_rate string
	// Contract number (as written on paper)
	// in: body
	//
	// required: true
	Number string
	// Number of days to pay after business trip expense appruval
	// in: body
	//
	// required: true
	Payment_period_bt string
	// Number of days to pay after timecard appruval
	// in: body
	//
	// required: true
	Payment_period_service string
	// PSoW contract (Partner Statement of Work)
	// in: body
	//
	// required: true
	Psow []psow
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
	// Contract expiraton date (YYYY-MM-DD)
	// in: body
	//
	// required: true
	End_date string
	// Current contract status
	// in: body
	//
	// required: true
	Status string
	// Subcontractor company object
	// in: body
	//
	// required: true
	Subcontractor_company []company
	// Array of tasks assigned to contract
	// in: body
	//
	// required: true
	Tasks []task
	// Array of persons to approve timecards
	// in: body
	//
	// required: true
	Timecard_approvers []approver
	// Reporting period (month or week)
	// in: body
	//
	// required: true
	Timecard_period string
	// Baseline of working day. Overtime calculaed based on this number.
	// in: body
	//
	// required: true
	Working_hours_per_day string
}

// Returns list of SoW (Statement of Works)
type sowServerResponse struct {
	// Result status
	// in: body
	Result string
	// Array of SoW
	// in: body
	Sow []sow
}

// Returns list of SoW (Statement of Works)
// swagger:response sowServerResponseWrapper
type sowServerResponseWrapper struct {
	// Result status
	// in: body
	Body sowServerResponse
}
