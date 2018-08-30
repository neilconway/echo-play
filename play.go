package main

import "github.com/labstack/echo"

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

func main() {
	m := &Master{}
	h := &Handler{
		master: m,
	}

	e := echo.New()
	e.GET("/agents", h.getAgents)
	e.GET("/agents/:agent_id", h.getAgent)
	e.GET("/agents/:agent_id/slots/:slot_id", h.getAgentSlot)

	e.Logger.Fatal(e.Start(":8080"))
}
