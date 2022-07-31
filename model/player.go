package model

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Name     string   // プレイヤー名
	Position Position // プレイヤーの位置情報
	Image    *ebiten.Image
}

// コンストラクタ
func NewPlayer(name string) *Player {
	img, _, err := ebitenutil.NewImageFromFile("assets/character/yuusya.png")
	if err != nil {
		log.Println(err)
	}

	position := NewPosition()
	return &Player{
		Name:     name,
		Position: *position,
		Image:    img,
	}
}
