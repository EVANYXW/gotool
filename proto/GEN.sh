#!/bin/sh
protoc  --go_out="../msg" --plugin=grpc:. Common.proto
protoc  --go_out="../msg" --plugin=grpc:. GMDefine.proto
protoc  --go_out="../msg" --plugin=grpc:. GM.proto
protoc  --go_out="../msg" --plugin=grpc:. LobbySystem.proto

protoc  --go_out="../msg" --plugin=grpc:. PlayerCenter.proto
protoc  --go_out="../msg" --plugin=grpc:. ServerDefine.proto
protoc  --go_out="../msg" --plugin=grpc:. PlayerCenter.proto
protoc  --go_out="../msg" --plugin=grpc:. Coon.proto
protoc  --go_out="../msg" --plugin=grpc:. Event.proto
protoc  --go_out="../msg" --plugin=grpc:. Game.proto
protoc  --go_out="../msg" --plugin=grpc:. Game_Slots.proto
protoc  --go_out="../msg" --plugin=grpc:. Lobby.proto


protoc  --go_out="../msg" --plugin=grpc:. RPC_LogService.rpc
protoc  --go_out="../msg" --plugin=grpc:. RPC_VipLoungeService.rpc
protoc  --go_out="../msg" --plugin=grpc:. RPC_MailService.rpc
