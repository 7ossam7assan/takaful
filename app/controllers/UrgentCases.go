package controllers

import (
	"Test/app/repositories"
	"github.com/revel/revel"
)

type UrgentCasesController struct {
	*revel.Controller
}

func (c UrgentCasesController) Index() revel.Result {
	UrgentCasesData := repositories.FindAll()
	data := make(map[string]interface{})
	data["data"] = UrgentCasesData
	return c.RenderJSON(data)
}
