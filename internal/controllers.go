package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping Ping API
// @summary Ping API
// @description In the realm of code where data flows,
// @description A whisper travels where the network goes.
// @description With a simple command, a heartbeat is sent,
// @description To check the connection, to see where it went.
// @description
// @description "Ping," it calls, like a friendly chime,
// @description A request for response, a dance through time.
// @description From client to server, a message takes flight,
// @description In packets it journeys, through day and through night.
// @description
// @description "Are you there?" it asks, with a digital sigh,
// @description A promise of data, a link to the sky.
// @description The echo returns, a confirmation so sweet,
// @description A symphony of bytes, where two systems meet.
// @description
// @description In the world of APIs, where services blend,
// @description Ping is the guardian, the reliable friend.
// @description It measures the distance, the latency's grace,
// @description A pulse of the network, a test of the space.
// @description
// @description So here’s to the ping, in its silent decree,
// @description A bridge in the ether, connecting you and me.
// @description With each little packet, a story unfolds,
// @description In the language of tech, where the future beholds.
// @tags base
// @produce json
// @success 200 {string} pong
// @router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
