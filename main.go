package main

import (
	_ "embed"
	_ "image/png"
	"log"

	// "github.com/golang/freetype/truetype"
	// "golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/shiiba-t/battle-royale/application"
	"github.com/shiiba-t/battle-royale/model"
)

// go:embed assets/fonts/x12y20pxScanLine.ttf
// var sampleBytes []byte

// var (
// 	gamerFontS font.Face
// )

//ゲーム全体に必要なデータを格納
type Game struct {
	Map *model.Map
}

// func init() {
// 	tt, err := truetype.Parse(sampleBytes)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	const dpi = 72
// 	gamerFontS = truetype.NewFace(tt, &truetype.Options{
// 		Size:    20,
// 		DPI:     dpi,
// 		Hinting: font.HintingFull,
// 	})
// }

// ゲーム構造体のコンストラクタ
func NewGame() *Game {
	g := &Game{}
	g.Map = application.CreateMap()
	return g
}

//Update is called each tic.
func (g *Game) Update() error {
	return nil
}

//Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	application.DrawMap(g.Map, screen)

	// text.Draw(screen, fmt.Sprintf("ABC"), gamerFontS, 0, 0, color.White)
	// text.Draw(screen, fmt.Sprintf("DEF"), gamerFontS, 0, 32, color.White)
	// text.Draw(screen, fmt.Sprintf("GHI"), gamerFontS, 32, 0, color.White)
	// text.Draw(screen, fmt.Sprintf("JKL"), gamerFontS, 32, 32, color.White)
}

//Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return 960, 480
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(960, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("SRPG")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
