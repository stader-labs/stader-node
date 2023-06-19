/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package math

import (
	"math"
)

// Round a float64 down to a number of places
func RoundDown(val float64, places int) float64 {
	return math.Floor(val*math.Pow10(places)) / math.Pow10(places)
}

// Round a float64 up to a number of places
func RoundUp(val float64, places int) float64 {
	return math.Ceil(val*math.Pow10(places)) / math.Pow10(places)
}
