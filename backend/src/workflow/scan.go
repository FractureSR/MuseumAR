package workflow

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"io"
	"strings"

	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
	"MuseumAR_Backend/infrastruture/webService"
)

type exhibitInfo struct {
	Name        string
	Description string
}

/*
	logic of Scan hanlder
	1 receive pictures from client
	2 do necessary transformation
	3 use Baidu picture matching api
	4 return results
*/

func Scan(ctx iris.Context) {
	//get the picture
	file, _, err := ctx.FormFile("upload")
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	//change it into bytes
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	//encode it into base64
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	//call recognize web service to cut the picture
	cutParams, err := webService.BaiduImageApi_object_detect(encoded)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	out, err := cut(buf, cutParams.Left, cutParams.Width, cutParams.Top, cutParams.Height)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	if buf, ok := out.(*bytes.Buffer); ok {
		encoded = base64.StdEncoding.EncodeToString(buf.Bytes())
	} else {
		ctx.StatusCode(500)
		return
	}

	//call search web service to find the exhibit
	id := ctx.URLParam("museum")
	e_id, err := webService.BaiduImageApi_same_hq_search(encoded, id)
	if err != nil {
		ctx.StatusCode(500)
		return
	}

	//get e_id from database
	db := database.Get()
	rows, err := db.Query([]string{"e_name", "e_description"}, []string{"exhibit"},
		"t_m_id="+"'"+id+"' and e_id="+"'"+e_id+"'", nil)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	var e exhibitInfo
	for rows.Next() {
		err = rows.Scan(&e.Name, &e.Description)
		if err != nil {
			ctx.StatusCode(500)
			return
		}
	}

	ctx.JSON(e)
}

func cut(in io.Reader, left int, width int, top int, height int) (io.Writer, error) {
	out := bytes.NewBuffer(nil)
	raw, format, err := image.Decode(in)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(format) {
	case "jpg":
	case "jpeg":
		img := raw.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(left, top, left+width, top-height)).(*image.YCbCr)
		err = jpeg.Encode(out, subImg, &jpeg.Options{100})
		if err != nil {
			return nil, err
		}
		return out, nil
	case "png":
		if img, ok := raw.(*image.RGBA); ok {
			subImg := img.SubImage(image.Rect(left, top, left+width, top-height)).(*image.RGBA)
			err = jpeg.Encode(out, subImg, &jpeg.Options{100})
			if err != nil {
				return nil, err
			}
			return out, nil
		} else if img, ok := raw.(*image.NRGBA); ok {
			subImg := img.SubImage(image.Rect(left, top, left+width, top-height)).(*image.NRGBA)
			err = jpeg.Encode(out, subImg, &jpeg.Options{100})
			if err != nil {
				return nil, err
			}
			return out, nil
		}
	}
	return nil, errors.New("Invalid picture format.")
}
