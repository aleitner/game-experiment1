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
		entity := NewTile((i%xNum)*tileSize, (i/xNum)*tileSize, id, sheets["tiles"])
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
	if ebiten.IsKeyPressed(ebiten.KeyUp) && r.player.currentActions["attacking"] == nil {
		r.player.direction = UP
		playerProposedMovement.y -= r.player.speed
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) && r.player.currentActions["attacking"] == nil {
		r.player.direction = DOWN
		playerProposedMovement.y += r.player.speed
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && r.player.currentActions["attacking"] == nil {
		r.player.direction = LEFT
		playerProposedMovement.x -= r.player.speed
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) && r.player.currentActions["attacking"] == nil {
		r.player.direction = RIGHT
		playerProposedMovement.x += r.player.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && r.player.currentActions["attacking"] == nil {
		r.player.currentActions["attacking"] = &FrameSpan{
			currentFrame: 0,
			endFrame:     5,
		}
		r.entities = append(r.entities, Swing(r.player))
	}

	r.player.hurtbox.x += playerProposedMovement.x
	r.player.hurtbox.y += playerProposedMovement.y

	// Remove entities with health of 0
	entities := r.entities[:0]
	for _, entity := range r.entities {
		for _, entity2 := range r.entities {
			if entity.Hits(entity2) && entity2.vulnerable {
				entity2.health--
			}

			if entity2.Hits(entity) && entity.vulnerable {
				entity.health--
			}
		}

		if r.player.Touches(entity) {
			r.player.hurtbox.x -= playerProposedMovement.x
			r.player.hurtbox.y -= playerProposedMovement.y
		}

		if r.player.Hits(entity) {
			entity.health--
		}

		if entity.Hits(r.player) {
			r.player.health--
		}

		if entity.health > 0 && ((entity.FrameSpan != nil && entity.currentFrame < entity.endFrame) || entity.FrameSpan == nil) {
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
		screen.DrawImage(sheets["tiles"].SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	}

	return nil
}
