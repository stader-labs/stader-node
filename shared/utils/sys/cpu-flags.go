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
package sys

import (
	"runtime"
	"sort"

	"github.com/klauspost/cpuid/v2"
)

// Returns the CPU features that are required to run "modern" images but are not present on the node's CPU
func GetMissingModernCpuFeatures() []string {
	var features map[cpuid.FeatureID]string
	switch runtime.GOARCH {
	case "amd64":
		features = map[cpuid.FeatureID]string{
			cpuid.ADX: "adx",
			//cpuid.AESNI: "aes",
			cpuid.AVX:   "avx",
			cpuid.AVX2:  "avx2",
			cpuid.BMI1:  "bmi1",
			cpuid.BMI2:  "bmi2",
			cpuid.CLMUL: "clmul",
			cpuid.MMX:   "mmx",
			cpuid.SSE:   "sse",
			cpuid.SSE2:  "sse2",
			cpuid.SSSE3: "ssse3",
			cpuid.SSE4:  "sse4.1",
			cpuid.SSE42: "sse4.2",
		}
	default:
		features = map[cpuid.FeatureID]string{}
	}

	unsupportedFeatures := []string{}
	for feature, name := range features {
		if !cpuid.CPU.Supports(feature) {
			unsupportedFeatures = append(unsupportedFeatures, name)
		}
	}

	sort.Strings(unsupportedFeatures)
	return unsupportedFeatures
}
