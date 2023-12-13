package sd

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

const (
	SDFloatStringEqualityThreshold = 1e-5 // 0.00001
)

var SDWeiEqualityThreshold = eth.EthToWei(SDFloatStringEqualityThreshold)

func almostEqual(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= SDFloatStringEqualityThreshold
}

func WeiAlmostEqual(lhs, rhs *big.Int) bool {
	diversity := new(big.Int).Sub(lhs, rhs)

	return diversity.CmpAbs(SDWeiEqualityThreshold) <= 0
}

func PromptChooseSDWithMaxMin(msg, errMsg string, min, max float64) (float64, error) {
	var utilityAmountFloat float64

	var errParse error

	for {
		s := cliutils.Prompt(
			msg,
			`^[0-9]\d*(\.\d+)?$`,
			errMsg)

		utilityAmountFloat, errParse = strconv.ParseFloat(s, 64)
		if errParse != nil {
			fmt.Println(errMsg)
			continue
		}

		if almostEqual(utilityAmountFloat, min) {
			utilityAmountFloat = min
		}

		if almostEqual(utilityAmountFloat, max) {
			utilityAmountFloat = max
		}

		if utilityAmountFloat < min || utilityAmountFloat > max {
			fmt.Println(errMsg)
			continue
		}

		break
	}

	return utilityAmountFloat, errParse
}
