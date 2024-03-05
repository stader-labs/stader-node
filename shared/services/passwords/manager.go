/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.5.0]

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
package passwords

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Config
const (
	MinPasswordLength = 12
	FileMode          = 0600
)

// Password manager
type PasswordManager struct {
	passwordPath string
}

// Create new password manager
func NewPasswordManager(passwordPath string) *PasswordManager {
	return &PasswordManager{
		passwordPath: passwordPath,
	}
}

// Check if the password has been set
func (pm *PasswordManager) IsPasswordSet() bool {
	_, err := ioutil.ReadFile(pm.passwordPath)
	return (err == nil)
}

// Get the password
func (pm *PasswordManager) GetPassword() (string, error) {

	// Read from disk
	password, err := ioutil.ReadFile(pm.passwordPath)
	if err != nil {
		return "", fmt.Errorf("Could not read password from disk: %w", err)
	}

	// Return
	return string(password), nil

}

// Set the password
func (pm *PasswordManager) SetPassword(password string) error {

	// Check password is not set
	if pm.IsPasswordSet() {
		return errors.New("Password is already set")
	}

	// Check password length
	if len(password) < MinPasswordLength {
		return fmt.Errorf("Password must be at least %d characters long", MinPasswordLength)
	}

	// Write to disk
	if err := ioutil.WriteFile(pm.passwordPath, []byte(password), FileMode); err != nil {
		return fmt.Errorf("Could not write password to disk: %w", err)
	}

	// Return
	return nil

}

// Delete the password
func (pm *PasswordManager) DeletePassword() error {

	// Check if it exists
	_, err := os.Stat(pm.passwordPath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return fmt.Errorf("error checking password file path: %w", err)
	}

	// Delete it
	err = os.Remove(pm.passwordPath)
	return err

}
