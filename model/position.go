package model

// (x,y)座標を格納する
// 左上が(0,0)
type Position struct {
	X int // x座標
	Y int // y座標
}

// コンストラクタ
func NewPosition() *Position {
	return &Position{
		X: 0,
		Y: 0,
	}
}
