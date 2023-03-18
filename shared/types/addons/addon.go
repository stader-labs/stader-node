package addons

// ROCKETPOOL-OWNED

import (
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
)

// Interface for Smartnode addons
type SmartnodeAddon interface {
	GetName() string
	GetDescription() string
	GetConfig() cfgtypes.Config
	GetContainerName() string
	GetContainerTag() string
	GetEnabledParameter() *cfgtypes.Parameter
	UpdateEnvVars(envVars map[string]string) error
}
