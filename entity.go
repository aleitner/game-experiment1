package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Entity struct {
	health             int
	img                *ebiten.Image
	w                  int
	h                  int
	x                  int
	y                  int
	v                  int
	facing             Direction
	isAttacking        bool
	attackFrame        int
	isTemporary        bool
	temporaryFrames    int
	maxTemporaryFrames int
	disappearsOnHit    bool
}

func NewPlayer(x, y int) *Entity {
	w := 9
	h := 12
	square, _ := ebiten.NewImage(w, h, ebiten.FilterNearest)

	// Fill the square with the white color
	square.Fill(color.RGBA{255, 0, 0, 255})

	return &Entity{
		health: 3,
		img:    square,
		x:      x,
		y:      y,
		v:      2,
		w:      w,
		h:      h,
	}
}

func NewTile(x, y, id int, tileSheet *ebiten.Image) *Entity {

	sx := (id % tileXNum) * tileSize
	sy := (id / tileXNum) * tileSize

	return &Entity{
		img:    tileSheet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image),
		x:      x,
		y:      y,
		health: 1,
		w:      tileSize,
		h:      tileSize,
	}
}

func NewEntity(id, x, y, w, h int) *Entity {
	square, _ := ebiten.NewImage(w, h, ebiten.FilterNearest)

	// Fill the square with the white color
	square.Fill(color.White)

	return &Entity{
		img: square,
		x:   x,
		y:   y,
		w:   w,
		h:   h,
	}
}

func (e *Entity) update(screen *ebiten.Image) error {

	if e.isTemporary {
		e.temporaryFrames++
	}

	if e.isAttacking {
		fmt.Println("Attacking... Frame: ", e.attackFrame)
		if e.attackFrame == 3 { // full animation
			e.isAttacking = false
			e.attackFrame = 0
		} else {
			e.attackFrame++
		}
	}

	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(float64(e.x), float64(e.y))

	// Draw the square image to the screen with an empty option
	screen.DrawImage(e.img, opts)

	return nil
}

func (e *Entity) CollidesWith(e2 *Entity) bool {

	c1 := e.y > e2.y+e2.h // e is above e2
	c2 := e2.y > e.y+e.h  // e2 is above e
	c3 := e2.x > e.x+e.w  // e2 is to the right of e
	c4 := e.x > e2.x+e2.w // e is to the right of e2

	return !c1 && !c2 && !c3 && !c4
}

func Swing(player *Entity) *Entity {
	var w int
	var h int
	var x int
	var y int

	switch player.facing {
	case UP:
		w = tileSize
		h = tileSize / 2
		x = player.x - ((w-player.w)/2 + 1)
		y = player.y - h - 1
	case DOWN:
		w = tileSize
		h = tileSize / 2
		x = player.x - ((w-player.w)/2 + 1)
		y = player.y + player.h + 1
	case LEFT:
		w = tileSize / 2
		h = tileSize
		x = player.x - w - 1
		y = player.y - ((h - player.h) / 2)
	case RIGHT:
		w = tileSize / 2
		h = tileSize
		x = player.x + player.w + 1
		y = player.y - ((h - player.h) / 2)
	}

	square, _ := ebiten.NewImage(w, h, ebiten.FilterNearest)

	// Fill the square with the white color
	square.Fill(color.RGBA{255, 0, 0, 255})

	return &Entity{
		img:                square,
		w:                  w,
		h:                  h,
		x:                  x,
		y:                  y,
		health:             1,
		isAttacking:        true,
		isTemporary:        true,
		maxTemporaryFrames: 3,
	}
}
