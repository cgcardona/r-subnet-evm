// (c) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package precompilebind

// tmplSourcePrecompileModuleGo is the Go precompiled module source template.
const tmplSourcePrecompileModuleGo = `
// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package {{.Package}}

import (
	"fmt"

	"github.com/cgcardona/r-subnet-evm/precompile/precompileconfig"
	"github.com/cgcardona/r-subnet-evm/precompile/contract"
	"github.com/cgcardona/r-subnet-evm/precompile/modules"

	"github.com/ethereum/go-ethereum/common"
)

var _ contract.Configurator = &configurator{}

// ConfigKey is the key used in json config files to specify this precompile precompileconfig.
// must be unique across all precompiles.
const ConfigKey = "{{decapitalise .Contract.Type}}Config"

// ContractAddress is the defined address of the precompile contract.
// This should be unique across all precompile contracts.
// See precompile/registry/registry.go for registered precompile contracts and more information.
var ContractAddress = common.HexToAddress("{ASUITABLEHEXADDRESS}") // SET A SUITABLE HEX ADDRESS HERE

// Module is the precompile module. It is used to register the precompile contract.
var Module = modules.Module{
	ConfigKey:    ConfigKey,
	Address:      ContractAddress,
	Contract:     {{.Contract.Type}}Precompile,
	Configurator: &configurator{},
}

type configurator struct{}

func init() {
	// Register the precompile module.
	// Each precompile contract registers itself through [RegisterModule] function.
	if err := modules.RegisterModule(Module); err != nil {
		panic(err)
	}
}

// MakeConfig returns a new precompile config instance.
// This is required for Marshal/Unmarshal the precompile config.
func (*configurator) MakeConfig() precompileconfig.Config {
	return new(Config)
}

// Configure configures [state] with the given [cfg] precompileconfig.
// This function is called by the EVM once per precompile contract activation.
// You can use this function to set up your precompile contract's initial state,
// by using the [cfg] config and [state] stateDB.
func (*configurator) Configure(chainConfig contract.ChainConfig, cfg precompileconfig.Config, state contract.StateDB, _ contract.BlockContext) error {
	config, ok := cfg.(*Config)
	if !ok {
		return fmt.Errorf("incorrect config %T: %v", config, config)
	}
	// CUSTOM CODE STARTS HERE
	{{- if .Contract.AllowList}}
	// AllowList is activated for this precompile. Configuring allowlist addresses here.
	return config.AllowListConfig.Configure(state, ContractAddress)
	{{- else}}
	return nil
	{{- end}}
}
`
