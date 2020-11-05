package count

import (
	sdk "github.com/irisnet/irishub-sdk-go"

	biftypes "github.com/SegueII/bifrost-1/types"
)

func CountRecordTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	CountRecordTask1(client, participants)
}

func CountRecordTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
}
