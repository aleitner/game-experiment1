package main

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

type Room struct {
	player     *Entity
	entities   []*Entity
	background []int
	foreground []int
}

func NewRoom(player *Entity, bg, mg, fg []int) *Room {
	room := &Room{
		player:     player,
		background: bg,
		foreground: fg,
	}

	for i, id := range mg {
		if id == 0 {
			continue
		}
		entity := NewTile((i%xNum)*tileSize, (i/xNum)*tileSize, id, tilesImage)
		room.entities = append(room.entities, entity)
	}

	return room
}

func (r *Room) update(screen *ebiten.Image) error {
	if err := r.drawBackground(screen); err != nil {
		return err
	}

	playerProposedMovement := &Movement{0, 0}

	// When the "up arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) && !r.player.isAttacking {
		r.player.facing = UP
		playerProposedMovement.y -= r.player.v
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) && !r.player.isAttacking {
		r.player.facing = DOWN
		playerProposedMovement.y += r.player.v
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && !r.player.isAttacking {
		r.player.facing = LEFT
		playerProposedMovement.x -= r.player.v
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) && !r.player.isAttacking {
		r.player.facing = RIGHT
		playerProposedMovement.x += r.player.v
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && !r.player.isAttacking {
		r.player.isAttacking = true
		r.entities = append(r.entities, Swing(r.player))
	}

	r.player.x += playerProposedMovement.x
	r.player.y += playerProposedMovement.y

	// Remove entities with health of 0
	entities := r.entities[:0]
	for _, entity := range r.entities {
		for _, entity2 := range r.entities {

			if entity.CollidesWith(entity2) && entity.isAttacking {
				entity2.health -= 1
			}
		}

		if r.player.CollidesWith(entity) {
			r.player.x -= playerProposedMovement.x
			r.player.y -= playerProposedMovement.y
		}

		if entity.health > 0 || entity.temporaryFrames < entity.maxTemporaryFrames{
			entities = append(entities, entity)
		}
	}
	r.entities = entities

	for _, entity := range r.entities {
		if err := entity.update(screen); err != nil {
			return err
		}
	}

	if err := r.player.update(screen); err != nil {
		return err
	}

	if err := r.drawForeground(screen); err != nil {
		return err
	}

	return nil
}

func (r *Room) drawBackground(screen *ebiten.Image) error {
	return r.drawEnvironment(r.background, screen)
}

func (r *Room) drawForeground(screen *ebiten.Image) error {
	return r.drawEnvironment(r.foreground, screen)
}

func (r *Room) drawEnvironment(tiles []int, screen *ebiten.Image) error {
	for i, t := range tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%xNum)*tileSize), float64((i/xNum)*tileSize))

		sx := (t % tileXNum) * tileSize
		sy := (t / tileXNum) * tileSize
		screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	}

	return nil
}
