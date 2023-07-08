package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/gogo/protobuf/plugin/stringer"
	"image/color"
)

var pathToPlayerBlackPieces = []string{
	"./assets/blackpieces/player/king.svg",
	"./assets/blackpieces/player/queen.svg",
	"./assets/blackpieces/player/rook.svg",
	"./assets/blackpieces/player/bishop.svg",
	"./assets/blackpieces/player/pawn.svg",
	"./assets/blackpieces/player/knight.svg",
}

var pathToOpponentBlackPieces = []string{
	"./assets/blackpieces/opponent/king.svg",
	"./assets/blackpieces/opponent/queen.svg",
	"./assets/blackpieces/opponent/rook.svg",
	"./assets/blackpieces/opponent/bishop.svg",
	"./assets/blackpieces/opponent/pawn.svg",
	"./assets/blackpieces/opponent/knight.svg",
}

var pathToPlayerWhitePieces = []string{
	"./assets/whitepieces/player/king.svg",
	"./assets/whitepieces/player/queen.svg",
	"./assets/whitepieces/player/rook.svg",
	"./assets/whitepieces/player/bishop.svg",
	"./assets/whitepieces/player/pawn.svg",
	"./assets/whitepieces/player/knight.svg",
}

var pathToOpponentWhitePieces = []string{
	"./assets/whitepieces/opponent/king.svg",
	"./assets/whitepieces/opponent/queen.svg",
	"./assets/whitepieces/opponent/rook.svg",
	"./assets/whitepieces/opponent/bishop.svg",
	"./assets/whitepieces/opponent/pawn.svg",
	"./assets/whitepieces/opponent/knight.svg",
}

type Board struct {
	Grid [][]Piece
}

func (b Board) setBoard() Board {
	board := make([][]Piece, 8)
	for i := range board {
		board[i] = make([]Piece, 8)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 8; j++ {
			switch {
			case i == 0 && j == 0 || i == 0 && j == 7:
				rook := Rook{
					Image:  "./assets/whitepieces/player/rook.svg",
					Player: "white",
					Pos:    Position{X: i, Y: j},
				}
				board[i][j] = rook

			case i == 0 && j == 1 || i == 0 && j == 6:
				knight := Knight{
					Image: "./assets/whitepieces/player/knight.svg",
					Player: "white",
					Pos: Position{X: i, Y: j})
				}
				board[i][j] = knight
			case i == 0 && j == 2 || i == 0 && j == 5:
				bishop := Bishop{
					Image: "./assets/whitepieces/player/bishop.svg",
					Player: "white",
					Pos: Position{X: i, Y: j}
				}
				board[i][j] = bishop
			case i == 0 && j == 3:
				queen := Queen{
					Image: "./assets/whitepieces/player/queen.svg",
					Player: "white",
					Pos: Position{X: i, Y: j}
				}
				board[i][j] = bishop
			case i == 0 && j == 4:
				king := King{
					Image: "./assets/whitepieces/player/king.svg",
					Player: "white",
					Pos: Position{X: i, Y: j}
				}
				board[i][j] = king
			default:
				pawn := Pawn{
					Image: "./assets/whitepieces/player/pawn.svg",
					Player: "white",
					Pos: Position{X: i, Y: j}
				}
				board[i][j] = pawn
			}
		}
	}

}

type Piece interface {
	availableMoves(gameBoard Board) [][]Position
	move(gameBoard *Board)
}

type King struct {
	Image  string
	Player string
	Pos    Position
}

/*
func (k King) availableMoves(gameBoard Board) [][]Position {
	return [][]int
}
*/
func (k King) move(gameBoard *Board) {

}

type Queen struct {
	Image    string
	Player   string
	Opponent string
	Pos      Position
}

func (q Queen) availableMoves(gameBoard Board) [][]Position {

}

func (q Queen) move(gameBoard Board) {

}

type Rook struct {
	Image    string
	Player   string
	Opponent string
	Pos      Position
}

func (r Rook) availableMoves(gameBoard Board) [][]Position {

}

func (r Rook) move(gameBoard *Board) {

}

type Bishop struct {
	Image    string
	Player   string
	Opponent string
	Pos      Position
}

func (b Bishop) availableMoves(gameBoard Board) [][]Position {

}

type Knight struct {
	Image  string
	Player string
	Pos    Position
}

func (k Knight) availableMoves(gameBoard Board) [][]Position {

}

type Pawn struct {
	Image  string
	Player string
	Pos    Position
}

func (p Pawn) availableMoves(gameBoard Board) [][]Position {

}

type Position struct {
	X int
	Y int
}

type Move struct {
	Start Position
	End   Position
}

type GUIBoard struct {
	Grid [][]fyne.CanvasObject
	Rows int
	Cols int
}

func newGUIBoard(rows, cols int) GUIBoard {
	board := make([][]fyne.CanvasObject, rows)
	for i := range board {
		board[i] = make([]fyne.CanvasObject, cols)
	}
	return GUIBoard{
		Grid: board,
		Rows: rows,
		Cols: cols,
	}
}

func paintChessBoard(chessBoard *GUIBoard) {
	for i := 0; i < chessBoard.Rows; i++ {
		for j := 0; j < chessBoard.Cols; j++ {
			rect := canvas.NewRectangle(&color.RGBA{R: 150, G: 77, B: 55, A: 1})
			if (i+j)%2 == 0 {
				rect.FillColor = color.White
			}
			rect.Refresh()
			chessBoard.Grid[i][j] = rect
		}
	}
}

func layoutForChessboard(board GUIBoard) fyne.CanvasObject {
	grid := container.NewGridWithColumns(board.Cols)
	for _, row := range board.Grid {
		for _, cell := range row {
			grid.Add(cell)
		}
	}
	return grid
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Chessboard")

	guiBoard := newGUIBoard(8, 8)
	paintChessBoard(&guiBoard)

	content := layoutForChessboard(guiBoard)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
