package workflow

import (
	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
)

type museumListItem struct {
	ID   string
	Name string
}

/*
	the workflow of AllMuseums
	1 fetch information from the database
	2 return it in json format
*/
func AllMuseums(ctx iris.Context) {
	//fetch museum list from database
	db := database.Get()
	rows, err := db.Query([]string{"m_id", "m_name"}, []string{"museum"}, "", nil)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	var museumList []museumListItem
	for rows.Next() {
		var ml museumListItem
		err = rows.Scan(&ml.ID, &ml.Name)
		if err != nil {
			ctx.StatusCode(500)
			return
		}
		museumList = append(museumList, ml)
	}

	//return the result
	ctx.JSON(iris.Map{
		"Museums": museumList,
	})
}
