package main

import (
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
	img *ebiten.Image
	currentActions map[string]*Action

	tangible   bool
	vulnerable bool
	hurtbox    *Box
	health     int
	direction  Direction
	speed      int

	hitbox *Box

	belongsTo *Entity
}

// Action keeps track of frames for an action
type Action struct {
	currentFrame int
	endFrame     int
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
		speed:      2,
		hurtbox: &Box {
			x:      x,
			y:      y,
			w:      w,
			h:      h,
		},
		direction: DOWN,
		tangible: true,
		vulnerable: true,
		currentActions: make(map[string]*Action),
	}
}

func NewTile(x, y, id int, tileSheet *ebiten.Image) *Entity {

	sx := (id % tileXNum) * tileSize
	sy := (id / tileXNum) * tileSize

	return &Entity{
		img:    tileSheet.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image),
		health: 1,
		hurtbox: &Box{
			x: x,
			y: y,
			w: tileSize,
			h: tileSize,
		},
		tangible: false,
		vulnerable: false,
	}
}

func (e *Entity) update(screen *ebiten.Image) error {

	for key, val := range e.currentActions {
		if val.currentFrame == val.endFrame {
			delete(e.currentActions, key)
		} else {
			val.currentFrame++
		}
	}

	//if e.isAttacking {
	//	fmt.Println("Attacking... Frame: ", e.attackFrame)
	//	if e.attackFrame == 3 { // full animation
	//		e.isAttacking = false
	//		e.attackFrame = 0
	//	} else {
	//		e.attackFrame++
	//	}
	//}

	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(float64(e.hurtbox.x), float64(e.hurtbox.y))

	// Draw the square image to the screen with an empty option
	screen.DrawImage(e.img, opts)

	return nil
}


func (attacker *Entity) CollidesWith(e *Entity) bool {
	return attacker.hitbox.CollidesWith(e.hurtbox)
}

func Swing(player *Entity) *Entity {
	var w int
	var h int
	var x int
	var y int

	switch player.direction {
	case UP:
		w = tileSize
		h = tileSize / 2
		x = player.hurtbox.x - ((w-player.hurtbox.w)/2 + 1)
		y = player.hurtbox.y - h - 1
	case DOWN:
		w = tileSize
		h = tileSize / 2
		x = player.hurtbox.x - ((w-player.hurtbox.w)/2 + 1)
		y = player.hurtbox.y + player.hurtbox.h + 1
	case LEFT:
		w = tileSize / 2
		h = tileSize
		x = player.hurtbox.x - w - 1
		y = player.hurtbox.y - ((h - player.hurtbox.h) / 2)
	case RIGHT:
		w = tileSize / 2
		h = tileSize
		x = player.hurtbox.x + player.hurtbox.w + 1
		y = player.hurtbox.y - ((h - player.hurtbox.h) / 2)
	}

	square, _ := ebiten.NewImage(w, h, ebiten.FilterNearest)

	// Fill the square with the white color
	square.Fill(color.RGBA{255, 0, 0, 255})

	return &Entity{
		img:                square,
		hurtbox: &Box {
			x:      x,
			y:      y,
			w:      w,
			h:      h,
		},
		health: 1,
		vulnerable: false,
		tangible: true,
		direction: player.direction,
	}
}
