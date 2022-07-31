package model

type Player struct {
	Name    string   // プレイヤー名
	Postion Position // プレイヤーの位置情報
}

// コンストラクタ
func NewPlayer(name string) *Player {
	position := NewPosition()
	return &Player{
		Name:    name,
		Postion: *position,
	}
}

// プレイヤーの初期位置を決める関数
func (p *Player) InitializePlayerPosition() {

}
