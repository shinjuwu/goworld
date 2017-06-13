package consts

import "time"

// Tunable Options
const (
	// For Server & Gate
	DISPATCHER_CLIENT_PACKET_QUEUE_SIZE = 100 // packet queue size
	// For Server
	SERVER_TICK_INTERVAL = time.Millisecond * 10 // server tick interval => affects timer resolution
	SAVE_INTERVAL        = time.Second * 10
	//SAVE_INTERVAL      = time.Minute * 5 // Save interval of entities
	// For Storage
)

// Debug Options
const (
	DEBUG_PACKETS   = false
	DEBUG_SPACES    = false
	DEBUG_SAVE_LOAD = false
	DEBUG_CLIENTS   = true
)