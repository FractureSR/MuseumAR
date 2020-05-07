package workflow

import (
	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
)

type ticketListItem struct {
	Name   string
	Detail string
	Price  int
}

func AllTickets(ctx iris.Context) {
	db := database.Get()
	id := ctx.URLParam("museum")
	rows, err := db.Query([]string{"t_name", "t_explain", "t_price"},
		[]string{"ticket"}, "t_m_id="+"'"+id+"'", nil)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	var ticketList []ticketListItem
	for rows.Next() {
		var tl ticketListItem
		err = rows.Scan(&tl.Name, &tl.Detail, &tl.Price)
		if err != nil {
			ctx.StatusCode(500)
			return
		}
		ticketList = append(ticketList, tl)
	}

	//return the result
	ctx.JSON(iris.Map{
		"Tickets": ticketList,
	})
}
