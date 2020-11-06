package main

import (
	"encoding/json"
	"io/ioutil"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/types/store"

	"github.com/SegueII/bifrost-1/count"
	biftypes "github.com/SegueII/bifrost-1/types"
)

var (
	nodeURI  = "http://34.80.22.255:26657"
	grpcAddr = "34.80.22.255:9090"
	chainID  = "bifrost-1"
)

func main() {
	client := initClient()

	participants := getParticipants()

	// count.CountServiceTasks(client, participants)
	count.CountRecordTasks(client, participants)
	// count.CountNFTTasks(client, participants)
	// count.CountRandomTasks(client, participants)

	bz, _ := json.MarshalIndent(participants, "", "    ")
	if err := ioutil.WriteFile("result.json", bz, 0666); err != nil {
		panic(err)
	}
}

func initClient() sdk.IRISHUBClient {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	return sdk.NewIRISHUBClient(cfg)
}

func getParticipants() []*biftypes.Participant {
	bytes, err := ioutil.ReadFile("participants.json")
	if err != nil {
		panic(err)
	}

	var participants []*biftypes.Participant
	if err := json.Unmarshal(bytes, &participants); err != nil {
		panic(err)
	}

	return participants
}
