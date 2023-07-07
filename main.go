import(
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)


struct Board{
	Grid [][] interface{}
	Rows int
	Cols int
}


func newBoard(rows , cols int) Board{
	board := make([][] interface{}, rows) 
	for i := range board{
		board[i] = make([] interface{}, cols)
	}
	return Board{
		Grid : board,
		Rows : rows,
		Cols : cols
	}
}

func (b *Board) Get(rows, cols) interface{}{
	return  b[rows][cols]
}

func (b *Board) Set(rows, cols, value interface{}){
	b[rows][cols] = value
}


func (b *Board) FillBoard(value interface{}){
	for i := range b.Rows{
		for j := range b.Cols{
			b[i][j] = value
		}
	}
}

func paintChessBoard(chessBoard *Board){ // kunne gjort den her generisk med en funksjon som parameter for å farg brettet.
	for i:= 0; i < chessBoard.Rows; i++{
		for j := 0; j < chessBoard.Cols; j++{
			var color string
			if (i+j) % 2 == 0{
				color = "#FFFFFF"
			} else{
				color = "#000000"
			}
			board[i][j] = canvas.NewRectangle(&fyne.StaticResource{
				StaticName: "color",
				Content: []byte(color),
			})
		}
	}

}

func main{
	myApp := app.New()
	chessBoard := myApp.NewWindow("Chessboard")

	board := newBoard(8,8)

	paintChessBoard(&board)







}