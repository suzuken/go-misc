// Package main for ...
//
// Schemes: http, https
// Consumes:
// - application/json
// Version: 2.0
//
// swagger:meta
package main

import (
	"fmt"

	"github.com/go-openapi/strfmt"
)

type unexportedType struct {
	nothingToShow string
}

// ExportedType should not generated in swagger.json but not yet.
//
// see: https://github.com/go-swagger/go-swagger/issues/796
type ExportedType struct {
	NothingToShow string
}

// ImportUnexportedField has field of unexported type.
//
// swagger:model
type ImportUnexportedField struct {
	F              unexportedType
	ExportedString string
	any            string
}

// User represents the user for this application
//
// A user is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A user can have friends with whom they can share what they like.
//
// swagger:model
type User struct {
	// the id for this user
	//
	// required: true
	// min: 1
	ID int64 `json:"id"`

	// the name for this user
	// required: true
	// min length: 3
	Name string `json:"name"`

	// the email address for this user
	//
	// required: true
	Email strfmt.Email `json:"login"`

	// the friends for this user
	Friends []User `json:"friends"`
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// ServeAPI serves the API for this record store
func ServeAPI(host, basePath string, schemes []string) error {

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: User
	//       200: someResponse
	//       201: ImportUnexportedField
	//       422: validationError
	mountItem("GET", basePath+"/users", nil)
	return nil
}

func mountItem(method, path string, listeners interface{}) {}

func main() {
	fmt.Println("vim-go")
}
