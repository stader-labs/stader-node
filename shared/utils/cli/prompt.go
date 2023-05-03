/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta]

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
package cli

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Prompt for user input
func Prompt(initialPrompt string, expectedFormat string, incorrectFormatPrompt string) string {

	// Print initial prompt
	fmt.Println(initialPrompt)

	// Get valid user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(); !regexp.MustCompile(expectedFormat).MatchString(scanner.Text()); scanner.Scan() {
		fmt.Println("")
		fmt.Println(incorrectFormatPrompt)
	}
	fmt.Println("")

	// Return user input
	return scanner.Text()

}

// Prompt for confirmation
func Confirm(initialPrompt string) bool {
	response := Prompt(fmt.Sprintf("%s [y/n]", initialPrompt), "(?i)^(y|yes|n|no)$", "Please answer 'y' or 'n'")
	return (strings.ToLower(response[:1]) == "y")
}

// Prompt for user selection
func Select(initialPrompt string, options []string) (int, string) {

	// Get prompt
	prompt := initialPrompt
	for i, option := range options {
		prompt += fmt.Sprintf("\n%d: %s", (i + 1), option)
	}

	// Get expected response format
	optionNumbers := []string{}
	for i := range options {
		optionNumbers = append(optionNumbers, strconv.Itoa(i+1))
	}
	expectedFormat := fmt.Sprintf("^(%s)$", strings.Join(optionNumbers, "|"))

	// Prompt user
	response := Prompt(prompt, expectedFormat, "Please enter a number corresponding to an option")

	// Get selected option
	index, _ := strconv.Atoi(response)
	selectedIndex := index - 1
	selectedOption := options[selectedIndex]

	// Return
	return selectedIndex, selectedOption

}

// Prompts the user to verify that there is nobody looking over their shoulder before printing sensitive information.
func ConfirmSecureSession(warning string) bool {
	if !Confirm(fmt.Sprintf("%s%s%s\nAre you sure you want to continue?", colorYellow, warning, colorReset)) {
		fmt.Println("Cancelled.")
		return false
	}

	return true
}
