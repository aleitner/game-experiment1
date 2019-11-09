package main

type Box struct {
	w int
	h int
	x int
	y int
}

func (hitbox *Box) Touches(hurtbox *Box) bool {
	if hurtbox == nil || (hurtbox.h == 0 && hurtbox.w == 0) {
		return false
	}

	c1 := hitbox.y > hurtbox.y+hurtbox.h // e is above e2
	c2 := hurtbox.y > hitbox.y+hitbox.h  // e2 is above e
	c3 := hurtbox.x > hitbox.x+hitbox.w  // e2 is to the right of e
	c4 := hitbox.x > hurtbox.x+hurtbox.w // e is to the right of e2

	return !c1 && !c2 && !c3 && !c4
}
