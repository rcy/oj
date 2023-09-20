package chess

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"oj/handlers/layout"
	"oj/handlers/render"
	"oj/templatehelpers"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/notnil/chess"
)

type Square struct {
	SVGPiece string
	Selected bool
	Action   string
}

type Board [8][8]Square

var pageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/fun/chess/page.gohtml"))

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game := chess.NewGame()

	log.Println(game.Position().Board().Draw())

	board := gameBoard(game, nil)

	render.Execute(w, pageTemplate, struct {
		Layout     layout.Data
		Board      Board
		ValidMoves []*chess.Move
	}{
		Layout:     l,
		Board:      board,
		ValidMoves: game.ValidMoves(),
	})
}

func gameBoard(game *chess.Game, position *Position) Board {
	pos := game.Position()
	return squareMapToBoard(pos.Board().SquareMap(), position)
}

type Position struct {
	rank int
	file int
}

func squareMapToBoard(squareMap map[chess.Square]chess.Piece, selected *Position) Board {
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

	var board Board

	for i := 0; i < 64; i += 1 {
		piece := squareMap[chess.Square(i)]
		rank := 7 - i/8
		file := i % 8
		square := &board[rank][file]
		square.SVGPiece = svgPiece[piece]

		if selected != nil && selected.rank == rank && selected.file == file {
			square.Action = "/fun/chess/unselect"
			square.Selected = true
		} else {
			square.Action = fmt.Sprintf("/fun/chess/select/%d/%d", rank, file)
		}
	}

	return board
}

func Select(w http.ResponseWriter, r *http.Request) {
	rankStr := chi.URLParam(r, "rank")
	fileStr := chi.URLParam(r, "file")

	rank, _ := strconv.Atoi(rankStr)
	file, _ := strconv.Atoi(fileStr)

	game := chess.NewGame()
	moves := game.ValidMoves()
	board := gameBoard(game, &Position{rank, file})

	selectedSquare := chess.Square((7-rank)*8 + file)
	actuallyValidMoves := []*chess.Move{}
	for _, move := range moves {
		if move.S1() == selectedSquare {
			actuallyValidMoves = append(actuallyValidMoves, move)
		}
	}

	render.ExecuteNamed(w, pageTemplate, "board", struct {
		Board      Board
		ValidMoves []*chess.Move
	}{
		Board:      board,
		ValidMoves: actuallyValidMoves,
	})
}

func Unselect(w http.ResponseWriter, r *http.Request) {
	game := chess.NewGame()
	board := gameBoard(game, nil)

	render.ExecuteNamed(w, pageTemplate, "board", struct{ Board Board }{board})
}

func Move(w http.ResponseWriter, r *http.Request) {
	r1Str := chi.URLParam(r, "r1")
	f1Str := chi.URLParam(r, "f1")
	r2Str := chi.URLParam(r, "r2")
	f2Str := chi.URLParam(r, "f2")

	r1, _ := strconv.Atoi(r1Str)
	f1, _ := strconv.Atoi(f1Str)
	r2, _ := strconv.Atoi(r2Str)
	f2, _ := strconv.Atoi(f2Str)

	game := chess.NewGame()
	board := gameBoard(game, nil)

	board[r1][f1].Selected = true
	board[r2][f2].Selected = true

	render.ExecuteNamed(w, pageTemplate, "board", struct{ Board Board }{board})
}
