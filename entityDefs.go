package main

import "github.com/hajimehoshi/ebiten"

var (
	sheets map[string]*ebiten.Image
)

type animation struct {
	tileIDs  []int
	sheet    string
	interval float32
	loops    bool
}

type definition struct {
	animations map[string]animation
}

var (
	definitions = map[string]*definition{
		"player": &definition{
			animations: map[string]animation{
				"idle-left": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"idle-right": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"idle-down": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"idle-up": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"walk-down": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"walk-left": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"walk-right": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"walk-up": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"swing-up": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"swing-down": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"swing-left": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
				"swing-right": animation{
					tileIDs: []int{1, 2, 3, 4},
					sheet:   "player",
				},
			},
		},
	}
)
