package api

// type MinipoolStatusResponse struct {
// 	Status         string            `json:"status"`
// 	Error          string            `json:"error"`
// 	Minipools      []MinipoolDetails `json:"minipools"`
// 	LatestDelegate common.Address    `json:"latestDelegate"`
// }

// type ValidatorDetails struct {
// 	Exists      bool     `json:"exists"`
// 	Active      bool     `json:"active"`
// 	Index       uint64   `json:"index"`
// 	Balance     *big.Int `json:"balance"`
// 	NodeBalance *big.Int `json:"nodeBalance"`
// }

//	type CanExitMinipoolResponse struct {
//		Status        string `json:"status"`
//		Error         string `json:"error"`
//		CanExit       bool   `json:"canExit"`
//		InvalidStatus bool   `json:"invalidStatus"`
//	}
type ExitMessage struct {
	ValidatorIndex uint64 `json:"validatorIndex"`
	Signature      string `json:"signature"`
	Epoch          uint64 `json:"epoch"`
}

type ExitMessageResponse struct {
	message ExitMessage `json:"message"`
}

// type CanProcessWithdrawalResponse struct {
// 	Status        string             `json:"status"`
// 	Error         string             `json:"error"`
// 	CanWithdraw   bool               `json:"canWithdraw"`
// 	InvalidStatus bool               `json:"invalidStatus"`
// 	GasInfo       rocketpool.GasInfo `json:"gasInfo"`
// }
// type ProcessWithdrawalResponse struct {
// 	Status string      `json:"status"`
// 	Error  string      `json:"error"`
// 	TxHash common.Hash `json:"txHash"`
// }

// type CanProcessWithdrawalAndFinaliseResponse struct {
// 	Status        string             `json:"status"`
// 	Error         string             `json:"error"`
// 	CanWithdraw   bool               `json:"canWithdraw"`
// 	InvalidStatus bool               `json:"invalidStatus"`
// 	GasInfo       rocketpool.GasInfo `json:"gasInfo"`
// }
// type ProcessWithdrawalAndFinaliseResponse struct {
// 	Status string      `json:"status"`
// 	Error  string      `json:"error"`
// 	TxHash common.Hash `json:"txHash"`
// }
