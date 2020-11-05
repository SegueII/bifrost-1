module github.com/SegueII/bifrost-1

go 1.15

require (
	github.com/irisnet/irishub-sdk-go v0.0.0-20201020101416-95956cdc5bde
	github.com/tidwall/gjson v1.6.3
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/irishub-sdk-go => /Users/segue/Segue/Bianjie/IRISHUB/irishub-sdk-go
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)