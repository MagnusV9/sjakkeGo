struct Board{
	Grid [][] interface{}
	Rows int
	Cols int
}


func newBoard(rows , cols int) Grid{
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

func (b *Board) DrawBoard(){

}


func main{

}