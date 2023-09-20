package chess

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
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
	Dot      bool
}

type Board [8][8]Square

type GameBoard struct {
	Board      Board
	ValidMoves []*chess.Move
}

var pageTemplate = template.Must(template.New("layout.gohtml").Funcs(templatehelpers.FuncMap).ParseFiles(layout.File, "handlers/fun/chess/page.gohtml"))

var game *chess.Game = chess.NewGame()

func Page(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l, err := layout.FromContext(ctx)
	if err != nil {
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < 20; i += 1 {
		moves := game.ValidMoves()
		if len(moves) > 0 {
			move := moves[rand.Intn(len(moves))]
			game.Move(move)
		}
	}

	log.Println(game.Position().Board().Draw())

	gameBoard := gameBoard(game, nil)

	render.Execute(w, pageTemplate, struct {
		Layout    layout.Data
		GameBoard GameBoard
	}{
		Layout:    l,
		GameBoard: gameBoard,
	})
}

func gameBoard(game *chess.Game, position *Position) GameBoard {
	pos := game.Position()
	gb := squareMapToBoard(pos.Board().SquareMap(), position)

	moves := game.ValidMoves()
	if position == nil {
		gb.ValidMoves = moves
	} else {
		selectedSquare := chess.Square((7-position.rank)*8 + position.file)
		for _, move := range moves {
			if move.S1() == selectedSquare {
				gb.ValidMoves = append(gb.ValidMoves, move)
			}
		}

		for _, move := range gb.ValidMoves {
			target := move.S2()
			square := &gb.Board[7-target/8][target%8]
			square.Dot = true
			square.Action = "chess/move/" + move.String()
		}
	}

	return gb
}

type Position struct {
	rank int
	file int
}

func squareMapToBoard(squareMap map[chess.Square]chess.Piece, selected *Position) GameBoard {
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

	var gb GameBoard

	for i := 0; i < 64; i += 1 {
		piece := squareMap[chess.Square(i)]
		rank := 7 - i/8
		file := i % 8
		square := &gb.Board[rank][file]
		square.SVGPiece = svgPiece[piece]

		if selected != nil && selected.rank == rank && selected.file == file {
			square.Action = "/fun/chess/unselect"
			square.Selected = true
		} else {
			square.Action = fmt.Sprintf("/fun/chess/select/%d/%d", rank, file)
		}
	}

	return gb
}

func Select(w http.ResponseWriter, r *http.Request) {
	rankStr := chi.URLParam(r, "rank")
	fileStr := chi.URLParam(r, "file")

	rank, _ := strconv.Atoi(rankStr)
	file, _ := strconv.Atoi(fileStr)

	gb := gameBoard(game, &Position{rank, file})

	render.ExecuteNamed(w, pageTemplate, "board", struct {
		GameBoard GameBoard
	}{
		GameBoard: gb,
	})
}

func Unselect(w http.ResponseWriter, r *http.Request) {
	gb := gameBoard(game, nil)

	render.ExecuteNamed(w, pageTemplate, "board", struct{ GameBoard GameBoard }{gb})
}
