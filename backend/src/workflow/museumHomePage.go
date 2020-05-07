package workflow

import (
	"log"
	"time"

	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
)

type musemInfo struct {
	Name        string
	Description string
	Pictures    []string
}

type sectionInfo struct {
	Name        string
	Description string
	Pictures    []string
}

type noticeInfo struct {
	Title     string
	Content   string
	Signature string
	Time      time.Time
	Picture   string
	TopFlag   bool
}

type museumHomePageInfo struct {
	Museum   musemInfo
	Sections []sectionInfo
	Notices  []noticeInfo
}

/*
	the workflow of MuseumHomePage
	1 fetch introduction
	2 fetch sections
	3 fetch notices
*/
func MuseumHomePage(ctx iris.Context) {
	var m museumHomePageInfo
	db := database.Get()
	//fetch introduction
	id := ctx.URLParam("museum")
	rows, err := db.Query([]string{"m_name", "m_description", "m_picture1",
		"m_picture2", "m_picture3"}, []string{"museum"}, "m_id="+"'"+id+"'", nil)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	var s1, s2, s3 string
	for rows.Next() {
		err = rows.Scan(&m.Museum.Name, &m.Museum.Description, &s1, &s2, &s3)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
	}
	m.Museum.Pictures = append(m.Museum.Pictures, s1)
	if s2 != "" {
		m.Museum.Pictures = append(m.Museum.Pictures, s2)
	}
	if s3 != "" {
		m.Museum.Pictures = append(m.Museum.Pictures, s3)
	}

	//fetch sections
	rows, err = db.Query([]string{"s_name", "s_description", "s_picture1",
		"s_picture2"}, []string{"section"}, "section.s_m_id="+"'"+id+"'", nil)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	for rows.Next() {
		var s sectionInfo
		err = rows.Scan(&s.Name, &s.Description, &s1, &s2)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
		if s1 != "" {
			s.Pictures = append(s.Pictures, s1)
		}
		if s2 != "" {
			s.Pictures = append(s.Pictures, s2)
		}
		m.Sections = append(m.Sections, s)
	}

	//fetch notices
	rows, err = db.Query([]string{"n_title", "n_content", "n_signature",
		"n_time", "n_picture", "n_topflag"}, []string{"notice"},
		"n_m_id="+"'"+id+"'", nil)
	for rows.Next() {
		var n noticeInfo
		err = rows.Scan(&n.Title, &n.Content, &n.Signature, &n.Time,
			&n.Picture, &n.TopFlag)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
		m.Notices = append(m.Notices, n)
	}

	ctx.JSON(m)
}
