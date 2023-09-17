package chess

import (
	"html/template"
	"log"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/templatehelpers"

	"github.com/notnil/chess"
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

	squareMap := pos.Board().SquareMap()

	svgPiece := [13]string{
		"", // empty piece
		"/assets/chess/wK.svg",
		"/assets/chess/wQ.svg",
		"/assets/chess/wR.svg",
		"/assets/chess/wB.svg",
		"/assets/chess/wN.svg",
		"/assets/chess/wP.svg",
		"/assets/chess/bK.svg",
		"/assets/chess/bQ.svg",
		"/assets/chess/bR.svg",
		"/assets/chess/bB.svg",
		"/assets/chess/bN.svg",
		"/assets/chess/bP.svg",
	}

	var rows [8][8]string

	for i := 0; i < 64; i += 1 {
		piece := squareMap[chess.Square(i)]
		rows[i/8][i%8] = svgPiece[piece]
	}

	log.Println(rows)

	d := struct {
		Layout layout.Data
		Pieces [8][8]string
	}{
		Layout: l,
		Pieces: rows,
	}

	render.Execute(w, MyPageTemplate, d)
}
