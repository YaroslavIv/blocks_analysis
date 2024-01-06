package client

import (
	"client/async"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Client struct {
	async  async.Async
	router *gin.Engine
	topic  string
}

func Init(asyncAddr, topic string) *Client {
	gin.SetMode(gin.ReleaseMode)
	return &Client{
		async:  async.Init(async.RABBITMQ, asyncAddr, topic),
		router: gin.Default(),
		topic:  topic,
	}
}

func (client *Client) getTop(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, client.async.Get(client.async.Send()))
}

func (client *Client) Run(ginAddr string) error {
	client.router.GET(fmt.Sprintf("/%s", client.topic), client.getTop)
	return client.router.Run(ginAddr)
}
