package model

type Map struct {
	Tiles []Tile
}

func NewMap(tiles []Tile) *Map {
	return &Map{
		Tiles: tiles,
	}
}
