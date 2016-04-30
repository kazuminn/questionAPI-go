package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	type Quego struct {
		Question string
		Answer   string
	}

	type Questionslice struct {
		Questions []Quego
	}

	e.Get("/v1/questions.json", func(c echo.Context) error {
		var s Questionslice
		s.Questions = append(s.Questions, Quego{Question: "Shanghai_VPN", Answer: "127.0.0.1"})
		s.Questions = append(s.Questions, Quego{Question: "Beijing_VPN", Answer: "127.0.0.2"})
		json, _ := json.Marshal(&s)
		return c.JSON(http.StatusCreated, &json)
	})
	e.Run(standard.New(":1323"))
}
