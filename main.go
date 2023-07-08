package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

type King struct {
	Image  string
	Player string
}

func (k King) availableMoves(gameBoard Board) [][]int {
	return [][]int
}

type Queen struct {
	Image string
	Player string
}

func (q Queen) availableMoves(gameBoard Board) [][]int{
	
}

type Rook struct{
	Image string
	Player string
}

func (r Rook) availableMoves(gameBoard Board) [][]int{
	
}

type Bishop struct {
	Image string
	Player string
}

func (b Bishop) availableMoves(gameBoard Board) [][]int{
	
}

type Knight struct{
	Image string
	Player string
}

func (k Knight) availableMoves(gameBoard Board) [][]int{
	
}

type Pawn struct{
	Image string
	Player string
}

func (p Pawn) availableMoves(gameBoard Board) [][]int{
	
}

type Board struct {
	Grid [][]fyne.CanvasObject
	Rows int
	Cols int
}

func newBoard(rows, cols int) Board {
	board := make([][]fyne.CanvasObject, rows)
	for i := range board {
		board[i] = make([]fyne.CanvasObject, cols)
	}
	return Board{
		Grid: board,
		Rows: rows,
		Cols: cols,
	}
}

func paintChessBoard(chessBoard *Board) {
	for i := 0; i < chessBoard.Rows; i++ {
		for j := 0; j < chessBoard.Cols; j++ {
			rect := canvas.NewRectangle(color.Black)
			if (i+j)%2 == 0 {
				rect.FillColor = color.White
			} else {
				rect.FillColor = color.Black
			}
			rect.Refresh()
			chessBoard.Grid[i][j] = rect
		}
	}
}

func layoutForChessboard(board Board) fyne.CanvasObject {
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

	board := newBoard(8, 8)
	paintChessBoard(&board)

	content := layoutForChessboard(board)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
