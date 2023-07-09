package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"io/ioutil"
	"log"
)

func contains(board [][]Position, pos Position) bool {
	for _, row := range board {
		for _, position := range row {
			if position == pos {
				return true
			}
		}
	}
	return false
}

const (
	player   = "player"
	opponent = "opponent"
)

var pieces = []string{"king", "queen", "rook", "bishop", "pawn", "knight"}

func getPathToPieces(color, role string) []string {
	var paths []string

	for _, piece := range pieces {
		paths = append(paths, "./assets"+"/"+color+"/"+role+"/"+piece+".svg")
	}

	return paths
}

type Position struct {
	X, Y int
}

type Piece interface {
	AvailableMoves(gameBoard *Board) [][]Position
	Move(gameBoard *Board, pos Position)
	Image() string
}

type ChessPiece struct {
	Image  string
	Player string
	Pos    Position
}

type King struct {
	ChessPiece
}

func (k King) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (k King) Image() string {
	return k.ChessPiece.Image
}

func (k *King) Move(gameBoard *Board, pos Position) {

	if contains(k.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[k.Pos.X][k.Pos.Y] = nil
		k.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = k
	}
}

type Queen struct {
	ChessPiece
}

func (q Queen) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (q *Queen) Move(gameBoard *Board, pos Position) {
	if contains(q.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[q.Pos.X][q.Pos.Y] = nil
		q.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = q
	}
}

func (q Queen) Image() string {
	return q.ChessPiece.Image
}

type Rook struct {
	ChessPiece
}

func (r Rook) Image() string {
	return r.ChessPiece.Image
}

func (r Rook) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (r *Rook) Move(gameBoard *Board, pos Position) {
	if contains(r.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[r.Pos.X][r.Pos.Y] = nil
		r.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = r
	}
}

type Bishop struct {
	ChessPiece
}

func (b Bishop) Image() string {
	return b.ChessPiece.Image
}

func (b Bishop) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (b *Bishop) Move(gameBoard *Board, pos Position) {
	if contains(b.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[b.Pos.X][b.Pos.Y] = nil
		b.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = b
	}
}

type Knight struct {
	ChessPiece
}

func (k Knight) Image() string {
	return k.ChessPiece.Image
}

func (k Knight) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (k *Knight) Move(gameBoard *Board, pos Position) {
	if contains(k.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[k.Pos.X][k.Pos.Y] = nil
		k.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = k
	}
}

type Pawn struct {
	ChessPiece
}

func (p Pawn) Image() string {
	return p.ChessPiece.Image
}

func (p Pawn) AvailableMoves(gameBoard *Board) [][]Position {
	return nil
}

func (p *Pawn) Move(gameBoard *Board, pos Position) {
	if contains(p.AvailableMoves(gameBoard), pos) {
		gameBoard.Grid[p.Pos.X][p.Pos.Y] = nil
		p.Pos = pos
		gameBoard.Grid[pos.X][pos.Y] = p
	}
}

type Board struct {
	Grid [][]Piece
}

func (b Board) IsLegalMove(pos Position) bool {
	legal := len(b.Grid)
	return pos.X <= legal && pos.X >= 0 && pos.Y <= legal && pos.Y >= 0
}

func NewBoard() *Board {
	board := &Board{
		Grid: make([][]Piece, 8),
	}

	for i := range board.Grid {
		board.Grid[i] = make([]Piece, 8)
	}

	return board
}

func (b *Board) SetupBoard() {
	for i := 0; i < 2; i++ {
		b.SetupRow(i, "black", "opponent")
	}
	for i := 6; i < 8; i++ {
		b.SetupRow(i, "white", "player")
	}
}

func (b *Board) SetupRow(row int, color, role string) {
	piecePaths := getPathToPieces(color, role)
	for col := 0; col < 8; col++ {
		b.Grid[row][col] = &Pawn{ChessPiece{piecePaths[4], color, Position{X: row, Y: col}}}
	}
	switch {
	case row%7 == 0:
		b.Grid[row][0], b.Grid[row][7] = &Rook{ChessPiece{piecePaths[2], color, Position{X: row, Y: 0}}}, &Rook{ChessPiece{piecePaths[2], color, Position{X: row, Y: 7}}}
		b.Grid[row][1], b.Grid[row][6] = &Knight{ChessPiece{piecePaths[5], color, Position{X: row, Y: 1}}}, &Knight{ChessPiece{piecePaths[5], color, Position{X: row, Y: 6}}}
		b.Grid[row][2], b.Grid[row][5] = &Bishop{ChessPiece{piecePaths[3], color, Position{X: row, Y: 2}}}, &Bishop{ChessPiece{piecePaths[3], color, Position{X: row, Y: 5}}}
		b.Grid[row][3] = &Queen{ChessPiece{piecePaths[1], color, Position{X: row, Y: 3}}}
		b.Grid[row][4] = &King{ChessPiece{piecePaths[0], color, Position{X: row, Y: 4}}}
	}
}

type GUIBoard struct {
	Grid       [][]fyne.CanvasObject
	Rows, Cols int
}

func newGUIBoard(rows, cols int) *GUIBoard {
	board := &GUIBoard{
		Grid: make([][]fyne.CanvasObject, rows),
		Rows: rows,
		Cols: cols,
	}

	for i := range board.Grid {
		board.Grid[i] = make([]fyne.CanvasObject, cols)
	}

	return board
}

func layoutForChessboard(board *GUIBoard) fyne.CanvasObject {
	grid := container.NewGridWithColumns(board.Cols)
	for _, row := range board.Grid {
		for _, cell := range row {
			grid.Add(cell)
		}
	}
	return grid
}

func drawPiece(piece Piece) *canvas.Image {
	svgData, err := ioutil.ReadFile(piece.Image())
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	svgResource := fyne.NewStaticResource(piece.Image(), svgData)

	return canvas.NewImageFromResource(svgResource)
}

func paintChessBoard(chessBoard *GUIBoard, gameBoard *Board) {
	for i := 0; i < chessBoard.Rows; i++ {
		for j := 0; j < chessBoard.Cols; j++ {
			rect := canvas.NewRectangle(&color.RGBA{R: 150, G: 77, B: 55, A: 1})
			if (i+j)%2 == 0 {
				rect.FillColor = color.White
			}
			rect.Refresh()

			// create a container for the square
			square := container.NewMax(rect)

			// draw piece if it exists
			piece := gameBoard.Grid[i][j]
			if piece != nil {
				img := drawPiece(piece)
				img.FillMode = canvas.ImageFillContain
				square.Add(img)
			}

			chessBoard.Grid[i][j] = square
		}
	}
}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Chess")

	guiBoard := newGUIBoard(8, 8)

	gameBoard := NewBoard()
	gameBoard.SetupBoard()

	paintChessBoard(guiBoard, gameBoard)

	content := layoutForChessboard(guiBoard)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
