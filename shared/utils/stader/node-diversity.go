package stader

import (
	"encoding/json"
	"fmt"

	"github.com/stader-labs/stader-node/shared/services"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"github.com/urfave/cli"
)

func SendNodeDiversityResponseType(
	c *cli.Context,
	request *stader_backend.NodeDiversityRequest,
) (*stader_backend.NodeDiversityResponseType, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	res, err := net.MakePostRequest(config.StaderNode.GetNodeDiversityApi(), request)
	if err != nil {
		return nil, fmt.Errorf("request to GetNodeDiversityApi %w", err)
	}
	defer res.Body.Close()

	var resp stader_backend.NodeDiversityResponseType
	err = json.NewDecoder(res.Body).Decode(&resp)

	if err != nil {
		return nil, fmt.Errorf("decode NodeDiversityResponseType %w", err)
	}

	return &resp, nil
}
