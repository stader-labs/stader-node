/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/stader-labs/stader-node/shared/types/api"
)

// Print an API response
// response must be a pointer to a struct type with Error and Status string fields
func PrintResponse(response interface{}, responseError error) {

	// Check response type
	r := reflect.ValueOf(response)
	if !(r.Kind() == reflect.Ptr && r.Type().Elem().Kind() == reflect.Struct) {
		PrintErrorResponse(errors.New("Invalid API response"))
		return
	}

	// Create zero response value if nil
	if r.IsNil() {
		response = reflect.New(r.Type().Elem()).Interface()
		r = reflect.ValueOf(response)
	}

	// Get and check response fields
	sf := r.Elem().FieldByName("Status")
	ef := r.Elem().FieldByName("Error")
	if !(sf.IsValid() && sf.CanSet() && sf.Kind() == reflect.String && ef.IsValid() && ef.CanSet() && ef.Kind() == reflect.String) {
		PrintErrorResponse(errors.New("Invalid API response"))
		return
	}

	// Populate error
	if responseError != nil {
		ef.SetString(responseError.Error())
	}

	// Set status
	if ef.String() == "" {
		sf.SetString("success")
	} else {
		sf.SetString("error")
	}

	// Encode
	responseBytes, err := json.Marshal(response)
	if err != nil {
		PrintErrorResponse(fmt.Errorf("Could not encode API response: %w", err))
		return
	}

	// Print
	fmt.Println(string(responseBytes))

}

// Print an API error response
func PrintErrorResponse(err error) {
	PrintResponse(&api.APIResponse{}, err)
}
