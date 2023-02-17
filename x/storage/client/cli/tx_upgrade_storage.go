package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/jackalLabs/canine-chain/x/storage/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpgradeStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-storage [for-address] [duration] [bytes] [payment-denom]",
		Short: "Broadcast message upgrade-storage",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argForAddress := args[0]
			argDuration := args[1]
			argBytes := args[2]
			argPaymentDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpgradeStorage(
				clientCtx.GetFromAddress().String(),
				argForAddress,
				argDuration,
				argBytes,
				argPaymentDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}