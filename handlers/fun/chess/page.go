package chess

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/templatehelpers"
	"strings"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

var MyPageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/fun/chess/page.gohtml"))

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game := chess.NewGame()
	pos := game.Position()

	m := pos.Board().SquareMap()
	for i := 0; i < 64; i += 1 {
		fmt.Println(m[chess.Square(i)])
	}

	boardSVG := new(strings.Builder)
	perspective := image.Perspective(chess.White)
	yellow := color.RGBA{255, 0, 0, 1}
	mark := image.MarkSquares(yellow, chess.D2, chess.E4)
	white := color.RGBA{255, 255, 255, 1}
	gray := color.RGBA{120, 120, 120, 1}
	sqrs := image.SquareColors(white, gray)
	err = image.SVG(boardSVG, pos.Board(), perspective, mark, sqrs)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b64svg := base64.StdEncoding.EncodeToString([]byte(boardSVG.String()))

	d := struct {
		Layout   layout.Data
		BoardSrc template.HTMLAttr
	}{
		Layout:   l,
		BoardSrc: template.HTMLAttr(fmt.Sprintf(`src="data:image/svg+xml;base64,%s"`, b64svg)),
	}

	render.Execute(w, MyPageTemplate, d)
}
