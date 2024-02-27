package sd

import (
	"fmt"
	"math/big"
	"strconv"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
)

const (
	SDFloatStringEqualityThreshold = 0.1
)

var SDWeiEqualityThreshold = eth.EthToWei(SDFloatStringEqualityThreshold)

func WeiAlmostEqual(lhs, rhs *big.Int) bool {
	diversity := new(big.Int).Sub(lhs, rhs)

	return diversity.CmpAbs(SDWeiEqualityThreshold) <= 0
}

func PromptChooseSDWithMaxMin(msg, errMsg string, min, max *big.Int) (*big.Int, error) {
	var utilityAmountWei *big.Int

	var errParse error

	for {
		s := cliutils.Prompt(
			msg,
			`^[0-9]\d*(\.\d+)?$`,
			errMsg)

		var utilityAmountFloat float64
		utilityAmountFloat, errParse = strconv.ParseFloat(s, 64)

		if errParse != nil {
			fmt.Println(errMsg)
			continue
		}

		utilityAmountWei = eth.EthToWei(utilityAmountFloat)

		if utilityAmountWei.Cmp(min) < 0 || utilityAmountWei.Cmp(max) > 0 || utilityAmountWei.Cmp(big.NewInt(0)) == 0 {
			fmt.Println(errMsg)
			continue
		}

		break
	}

	return utilityAmountWei, errParse
}
