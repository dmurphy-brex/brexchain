package txnengine

import (
	"github.com/cosmonaut/brexchain/x/txnengine/keeper"
	"github.com/cosmonaut/brexchain/x/txnengine/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the transaction
	for _, elem := range genState.TransactionList {
		k.SetTransaction(ctx, elem)
	}

	// Set transaction count
	k.SetTransactionCount(ctx, genState.TransactionCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TransactionList = k.GetAllTransaction(ctx)
	genesis.TransactionCount = k.GetTransactionCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
