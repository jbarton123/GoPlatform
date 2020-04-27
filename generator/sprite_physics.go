package generator

import "github.com/faiface/pixel"

type SpritePhysics interface {
	Update(dt float64, ctrl pixel.Vec, platforms []Platform)
}

type spritePhysics struct {
	gravity   float64
	runSpeed  float64
	jumpSpeed float64
	rect      pixel.Rect
	vel       pixel.Vec
	ground    bool
}

func NewSpritePhysicsUpdater(gravity float64, runSpeed float64, jumpSpeed float64, rect pixel.Rect, vel pixel.Vec, ground bool) SpritePhysics {
	return &spritePhysics{
		gravity:   gravity,
		runSpeed:  runSpeed,
		jumpSpeed: jumpSpeed,
		rect:      rect,
		vel:       vel,
		ground:    ground,
	}
}

func (gp *spritePhysics) Update(dt float64, ctrl pixel.Vec, platforms []Platform) {
	// apply controls
	switch {
	case ctrl.X < 0:
		gp.vel.X = -gp.runSpeed
	case ctrl.X > 0:
		gp.vel.X = +gp.runSpeed
	default:
		gp.vel.X = 0
	}

	// apply gravity and velocity
	gp.vel.Y += gp.gravity * dt
	gp.rect = gp.rect.Moved(gp.vel.Scaled(dt))

	// check collisions against each platform
	gp.ground = false
	if gp.vel.Y <= 0 {
		for _, p := range platforms {
			if gp.rect.Max.X <= p.rect.Min.X || gp.rect.Min.X >= p.rect.Max.X {
				continue
			}
			if gp.rect.Min.Y > p.rect.Max.Y || gp.rect.Min.Y < p.rect.Max.Y+gp.vel.Y*dt {
				continue
			}
			gp.vel.Y = 0
			gp.rect = gp.rect.Moved(pixel.V(0, p.rect.Max.Y-gp.rect.Min.Y))
			gp.ground = true
		}
	}

	// jump if on the ground and the player wants to jump
	if gp.ground && ctrl.Y > 0 {
		gp.vel.Y = gp.jumpSpeed
	}
}
