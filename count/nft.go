package count

import (
	sdk "github.com/irisnet/irishub-sdk-go"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/SegueII/bifrost-1/types"
)

func CountNFTTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	CountNFTTask1(client, participants)
	CountNFTTask2(client, participants)
	CountNFTTask3(client, participants)
	CountNFTTask4(client, participants)
	CountNFTTask5(client, participants)
}

func CountNFTTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("issue_denom")),
		)

		txs, err := client.QueryTxs(builder, 1, 10000)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks[8] = true
		}
	}
}

func CountNFTTask2(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}

func CountNFTTask3(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}

func CountNFTTask4(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}

func CountNFTTask5(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}
