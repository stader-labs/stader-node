/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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
package hex

// Add a prefix to a hex string if not present
func AddPrefix(value string) string {
	if len(value) < 2 || value[0:2] != "0x" {
		return "0x" + value
	}
	return value
}

// Remove a prefix from a hex string if present
func RemovePrefix(value string) string {
	if len(value) >= 2 && value[0:2] == "0x" {
		return value[2:]
	}
	return value
}
