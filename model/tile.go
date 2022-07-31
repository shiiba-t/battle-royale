package model

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileType int

// マップのタイルの種類を定義
const (
	TileTypeGrassland TileType = iota // 草原
)

type Tile struct {
	Position Position      // 座標位置(x,y)
	Image    *ebiten.Image // 画像
	TileType TileType      // タイプ
}

// コンストラクタ
func NewTile(position Position) *Tile {
	img, _, err := ebitenutil.NewImageFromFile("assets/map/grasslands.png")
	if err != nil {
		log.Println(err)
	}
	return &Tile{
		Position: position,
		Image:    img,
		TileType: TileTypeGrassland,
	}
}
