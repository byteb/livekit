// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/protocol/auth"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config, keyProvider auth.KeyProvider, currentNode routing.LocalNode) (*LivekitServer, error) {
	client, err := createRedisClient(conf)
	if err != nil {
		return nil, err
	}
	roomStore := createStore(client)
	router := createRouter(client, currentNode)
	nodeSelector := nodeSelectorFromConfig(conf)
	notifier, err := createWebhookNotifier(conf, keyProvider)
	if err != nil {
		return nil, err
	}
	roomManager, err := NewRoomManager(roomStore, router, currentNode, nodeSelector, notifier, conf)
	if err != nil {
		return nil, err
	}
	roomService, err := NewRoomService(roomManager)
	if err != nil {
		return nil, err
	}
	recordingService := NewRecordingService(client)
	rtcService := NewRTCService(conf, roomManager, router, currentNode)
	server, err := NewTurnServer(conf, roomStore, currentNode)
	if err != nil {
		return nil, err
	}
	livekitServer, err := NewLivekitServer(conf, roomService, recordingService, rtcService, keyProvider, router, roomManager, server, currentNode)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil
}

func InitializeRouter(conf *config.Config, currentNode routing.LocalNode) (routing.Router, error) {
	client, err := createRedisClient(conf)
	if err != nil {
		return nil, err
	}
	router := createRouter(client, currentNode)
	return router, nil
}
