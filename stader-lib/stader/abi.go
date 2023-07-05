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
package stader

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Decode, decompress and parse a zlib-compressed, base64-encoded ABI
func DecodeAbi(abiEncoded string) (*abi.ABI, error) {

	// base64 decode
	abiCompressed, err := base64.StdEncoding.DecodeString(abiEncoded)
	if err != nil {
		return nil, fmt.Errorf("Could not decode base64 data: %w", err)
	}

	// zlib decompress
	byteReader := bytes.NewReader(abiCompressed)
	zlibReader, err := zlib.NewReader(byteReader)
	if err != nil {
		return nil, fmt.Errorf("Could not decompress zlib data: %w", err)
	}
	defer func() {
		_ = zlibReader.Close()
	}()

	// Parse ABI
	abiParsed, err := abi.JSON(zlibReader)
	if err != nil {
		return nil, fmt.Errorf("Could not parse JSON: %w", err)
	}

	// Return
	return &abiParsed, nil

}

// zlib-compress and base64-encode an ABI JSON string
func EncodeAbiStr(abiStr string) (string, error) {

	// zlib compress
	var abiCompressed bytes.Buffer
	zlibWriter := zlib.NewWriter(&abiCompressed)
	if _, err := zlibWriter.Write([]byte(abiStr)); err != nil {
		return "", fmt.Errorf("Could not zlib-compress ABI string: %w", err)
	}
	if err := zlibWriter.Flush(); err != nil {
		return "", fmt.Errorf("Could not zlib-compress ABI string: %w", err)
	}
	if err := zlibWriter.Close(); err != nil {
		return "", fmt.Errorf("Could not zlib-compress ABI string: %w", err)
	}

	// base64 encode & return
	return base64.StdEncoding.EncodeToString(abiCompressed.Bytes()), nil

}
