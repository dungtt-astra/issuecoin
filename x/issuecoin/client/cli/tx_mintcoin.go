package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"log"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"issuecoin/x/issuecoin/types"
)

var _ = strconv.Itoa(0)

func CmdMintcoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mintcoin [coin] [to]",
		Short: "Broadcast message mintcoin",
		Args:  cobra.ExactArgs(2),
		//RunE: func(cmd *cobra.Command, args []string) (err error) {
		//	argCoin := args[0]
		//	argTo := args[1]
		//
		//	clientCtx, err := client.GetClientTxContext(cmd)
		//	if err != nil {
		//		return err
		//	}
		//
		//	msg := types.NewMsgMintcoin(
		//		clientCtx.GetFromAddress().String(),
		//		argCoin,
		//		argTo,
		//	)
		//	if err := msg.ValidateBasic(); err != nil {
		//		return err
		//	}
		//	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		//},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//argCoin := args[0]
			argTo := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			decCoin, err := sdk.ParseDecCoin(args[0])
			if err != nil {
				return err
			}
			mintCoin, _ := sdk.NormalizeDecCoin(decCoin).TruncateDecimal()

			msg := types.NewMsgMintcoin(
				//clientCtx.GetFromAddress().String(),
				argTo,
				&mintCoin,
				argTo,
			)
			log.Println("start ValidateBasic() .. to:", argTo)
			if err := msg.ValidateBasic(); err != nil {
				log.Println(" ValidateBasic err: %v", err)
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
