package application

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shiiba-t/battle-royale/model"
)

// 新規プレイヤーを作成する
func CreatePlayer() *model.Player {
	newPlayer := model.NewPlayer("勇者")
	InitPlayerPosition(newPlayer)

	return newPlayer
}

// プレイヤーを描画する
func DrawPlayer(player *model.Player, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(player.Position.X*32), float64(player.Position.Y*32))
	screen.DrawImage(player.Image, op)
}

// プレイヤーの初期位置を決める関数
func InitPlayerPosition(player *model.Player) {
	player.Position = SelectTileFromMap()
}
