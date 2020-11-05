package count

import (
	sdk "github.com/irisnet/irishub-sdk-go"

	biftypes "github.com/SegueII/bifrost-1/types"
)

func CountRandomTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	CountRandomTask1(client, participants)
	CountRandomTask2(client, participants)
}

func CountRandomTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	// for _, participant := range participants {
	// 	builder := types.NewEventQueryBuilder().AddCondition(
	// 		types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
	// 	).AddCondition(
	// 		types.NewCond("message", "action").EQ(types.EventValue("request_rand")),
	// 	)

	// 	txs, err := client.QueryTxs(builder, 1, 10000)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for _, tx := range txs.Txs {
	// 		requestID, err := tx.Result.Events.GetValue("message", "request_id")
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		println(requestID)
	// 	}
	// }
}

func CountRandomTask2(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}
