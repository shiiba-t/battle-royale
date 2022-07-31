package application

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shiiba-t/battle-royale/model"
)

const MAP_WIDTH = 30  // マップ全体の横幅
const MAP_HEIGHT = 15 // マップ全体の縦幅

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

func DrawMap(m *model.Map, screen *ebiten.Image) {
	for _, v := range m.Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.Position.X*32), float64(v.Position.Y*32))
		screen.DrawImage(v.Image, op)
	}
}
