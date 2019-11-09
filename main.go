package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	tileSize = 16
	tileXNum = 25

	tileWidth  = 9
	tileHeight = 10

	screenWidth  = tileSize * tileWidth
	screenHeight = tileSize * tileHeight

	xNum = screenWidth / tileSize
)

var State *state

type state struct {
	cameraX     int
	cameraY     int
	CurrentRoom *Room
	Player      *Entity
}

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.

	tilesF, err := os.Open("resources/tiles.png")
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(tilesF)
	if err != nil {
		log.Fatal(err)
	}

	sheets = make(map[string]*ebiten.Image)
	sheets["tiles"], _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if err := State.CurrentRoom.update(screen); err != nil {
		return err
	}

	return nil
}

func main() {
	player := NewPlayer(0, 0)
	State = &state{
		Player:      player,
		CurrentRoom: NewRoom(player, room2Background, room2Middleground, room2Foreground),
	}

	if err := ebiten.Run(update, screenWidth, screenHeight, 4, "Field Test"); err != nil {
		log.Fatal(err)
	}
}

type Movement struct {
	x int
	y int
}
