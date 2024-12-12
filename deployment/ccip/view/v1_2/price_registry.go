package v1_2

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/deployment/common/view/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry_1_2_0"
)

type PriceRegistryView struct {
	types.ContractMetaData
	FeeTokens          []common.Address `json:"feeTokens"`
	StalenessThreshold string           `json:"stalenessThreshold"`
	Updaters           []common.Address `json:"updaters"`
}

func GeneratePriceRegistryView(pr *price_registry_1_2_0.PriceRegistry) (PriceRegistryView, error) {
	if pr == nil {
		return PriceRegistryView{}, fmt.Errorf("cannot generate view for nil PriceRegistry")
	}
	meta, err := types.NewContractMetaData(pr, pr.Address())
	if err != nil {
		return PriceRegistryView{}, fmt.Errorf("failed to generate contract metadata for PriceRegistry %s: %w", pr.Address(), err)
	}
	ft, err := pr.GetFeeTokens(nil)
	if err != nil {
		return PriceRegistryView{}, fmt.Errorf("failed to get fee tokens %s: %w", pr.Address(), err)
	}
	st, err := pr.GetStalenessThreshold(nil)
	if err != nil {
		return PriceRegistryView{}, fmt.Errorf("failed to get staleness threshold %s: %w", pr.Address(), err)
	}
	updaters, err := pr.GetPriceUpdaters(nil)
	if err != nil {
		return PriceRegistryView{}, fmt.Errorf("failed to get price updaters %s: %w", pr.Address(), err)
	}
	return PriceRegistryView{
		ContractMetaData:   meta,
		FeeTokens:          ft,
		StalenessThreshold: st.String(),
		Updaters:           updaters,
	}, nil
}