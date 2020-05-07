package workflow

import (
	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
)

type souvenirListItem struct {
	Name     string
	Detail   string
	Price    string
	Stock    string
	Pictures []string
}

func AllSouvenirs(ctx iris.Context) {
	db := database.Get()
	id := ctx.URLParam("museum")
	rows, err := db.Query([]string{"k_name", "k_detail", "k_price", "k_stock",
		"k_picture1", "k_picture2", "k_picture3"}, []string{"keepsakes"},
		"k_m_id="+"'"+id+"'", nil)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	var souvenirList []souvenirListItem
	for rows.Next() {
		var sl souvenirListItem
		var s1, s2, s3 string
		err = rows.Scan(&sl.Name, &sl.Detail, &sl.Price, &sl.Stock, &s1, &s2, &s3)
		if err != nil {
			ctx.StatusCode(500)
			return
		}
		sl.Pictures = append(sl.Pictures, s1)
		if s2 != "" {
			sl.Pictures = append(sl.Pictures, s2)
		}
		if s3 != "" {
			sl.Pictures = append(sl.Pictures, s3)
		}
		souvenirList = append(souvenirList, sl)
	}

	ctx.JSON(iris.Map{
		"Souvenirs": souvenirList,
	})
}
