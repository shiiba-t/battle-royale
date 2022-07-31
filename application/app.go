package application

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shiiba-t/battle-royale/model"
)

const MAP_WIDTH = 30  // マップ全体の横幅
const MAP_HEIGHT = 15 // マップ全体の縦幅

// 新規マップを生成する
func CreateMap() *model.Map {
	newMap := model.NewMap([]model.Tile{})
	tiles := make([]model.Tile, 0, MAP_WIDTH*MAP_HEIGHT)

	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			tile := *model.NewTile(model.Position{X: j, Y: i})
			tiles = append(tiles, tile)
		}
	}

	newMap.Tiles = tiles
	return newMap
}

// マップを描画する
func DrawMap(m *model.Map, screen *ebiten.Image) {
	for _, v := range m.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.Position.X*32), float64(v.Position.Y*32))
		screen.DrawImage(v.Image, op)
	}
}

// マップの中からランダムなタイルを１枚選択する
func SelectTileFromMap() model.Position {
	rand.Seed(time.Now().UnixNano())

	return model.Position{
		X: rand.Intn(MAP_WIDTH),
		Y: rand.Intn(MAP_HEIGHT),
	}
}
