package genesis

import (
	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	tmtypes "github.com/tendermint/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	crypto.RegisterAmino(cdc)
}

var TestnetGenesis = `{
  "genesis_time": "2018-06-09T00:00:00Z",
  "chain_id": "minter-test-network-11",
  "validators": [
    {
      "pub_key": {
        "type": "tendermint/PubKeyEd25519",
        "value": "qu4d3zD/VMkHFdkotWZS/FEb7Tci5Ylz6O+Ub12uOXk="
      },
      "power": "100",
      "name": ""
    }
  ],
  "app_state": {
    "first_validator_address": "Mxa93163fdf10724dc4785ff5cbfb9ac0b5949409f",
    "initial_balances": [
      {
        "address": "Mxa93163fdf10724dc4785ff5cbfb9ac0b5949409f",
        "balance": {
          "MNT": "10000000000000000000000000"
        }
      },
      {
        "address": "Mxfe60014a6e9ac91618f5d1cab3fd58cded61ee99",
        "balance": {
          "MNT": "10000000000000000000000000"
        }
      }
    ]
  },
  "app_hash": "0000000000000000000000000000000000000000000000000000000000000000"
}`

func GetTestnetGenesis() *tmtypes.GenesisDoc {
	genDoc := tmtypes.GenesisDoc{}
	cdc.UnmarshalJSON([]byte(TestnetGenesis), &genDoc)
	genDoc.ValidateAndComplete()

	return &genDoc
}

type AppState struct {
	FirstValidatorAddress types.Address `json:"first_validator_address"`
	InitialBalances       []Account     `json:"initial_balances"`
}

type Account struct {
	Address types.Address     `json:"address"`
	Balance map[string]string `json:"balance"`
}
