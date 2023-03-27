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
package json

import (
	"encoding/json"
	"fmt"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("%w\nUnable to Unmarshal JSON string %s", err, string(data))
	}

	return nil
}
