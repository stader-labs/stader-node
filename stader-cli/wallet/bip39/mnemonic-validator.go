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
package bip39

import (
	"errors"
	"sort"
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type MnemonicValidator struct {
	mnemonic []string
}

func Create(length int) *MnemonicValidator {

	if length <= 0 {
		return nil
	}

	out := &MnemonicValidator{}

	out.mnemonic = make([]string, 0, length)

	return out
}

func (mv *MnemonicValidator) AddWord(input string) error {
	wordList := bip39.GetWordList()

	idx := sort.SearchStrings(wordList, input)
	if idx >= len(wordList) {
		return errors.New("Invalid word")
	}

	if wordList[idx] != input && (len(input) < 4 || wordList[idx][:4] != input[:4]) {
		return errors.New("Invalid word")
	}

	mv.mnemonic = append(mv.mnemonic, wordList[idx])
	return nil
}

func (mv *MnemonicValidator) Filled() bool {
	return len(mv.mnemonic) == cap(mv.mnemonic)
}

func (mv *MnemonicValidator) Finalize() (string, error) {

	if mv.Filled() == false {
		return "", errors.New("Not enough words were entered.")
	}

	mnemonic := strings.Join(mv.mnemonic, " ")
	if bip39.IsMnemonicValid(mnemonic) {
		return mnemonic, nil
	}

	return "", errors.New("Invalid mnemonic")
}
