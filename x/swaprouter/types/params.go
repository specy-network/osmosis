package types

import (
	"fmt"

	appparams "github.com/osmosis-labs/osmosis/v13/app/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyPoolCreationFee     = []byte("PoolCreationFee")
	defaultPoolCreationFee = sdk.NewInt64Coin(appparams.BaseCoinUnit, 1000_000_000) // 1000 OSMO
)

// ParamTable for gamm module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(poolCreationFee sdk.Coins) Params {
	return Params{
		PoolCreationFee: poolCreationFee,
	}
}

// DefaultParams are the default swaprouter module parameters.
func DefaultParams() Params {
	return Params{
		PoolCreationFee: sdk.NewCoins(defaultPoolCreationFee),
	}
}

// validate params.
func (p Params) Validate() error {
	if err := validatePoolCreationFee(p.PoolCreationFee); err != nil {
		return err
	}

	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPoolCreationFee, &p.PoolCreationFee, validatePoolCreationFee),
	}
}

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Validate() != nil {
		return fmt.Errorf("invalid pool creation fee: %+v", i)
	}

	return nil
}
