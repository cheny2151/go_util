package entiry

import (
	"strconv"
)

type GradeCost struct {
	Id         int
	Grade      string
	CostId     int
	Remark     string
}

func (gc GradeCost) ToString() string {
	return "{\r\n"+"Id=" + strconv.Itoa(gc.Id) + "\r\n" + "Grade=" + gc.Grade + "\r\n" + "CostId=" + strconv.Itoa(gc.CostId) + "\r\n" + "Remark=" + gc.Remark+"\r\n}"
}
