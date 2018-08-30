package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

type Master struct {
}

type Handler struct {
	master *Master
}

func (h *Handler) getAgents(c echo.Context) error {
	return nil
}

func (h *Handler) getAgent(c echo.Context) error {
	return nil
}

func (h *Handler) getAgentSlot(c echo.Context) error {
	return nil
}

func (h *Handler) agentWebSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Read startup message.
	_, msg, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
	}

	return nil
}

func (h *Handler) trialWebSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	return nil
}

func main() {
	m := &Master{}
	h := &Handler{
		master: m,
	}

	e := echo.New()
	e.GET("/agents", h.getAgents)
	e.GET("/agents/:agent_id", h.getAgent)
	e.GET("/agents/:agent_id/slots/:slot_id", h.getAgentSlot)

	e.GET("/ws/agent", h.agentWebSocket)
	e.GET("/ws/trial", h.trialWebSocket)

	e.Logger.Fatal(e.Start(":8080"))
}
