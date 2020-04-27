package generator

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
)

type PlatformGenerator interface {
	Generate(imd *imdraw.IMDraw)
}

type Platform struct {
	rect  pixel.Rect
	color color.Color
}

func NewPlatformGenerator(rect pixel.Rect, color color.Color) PlatformGenerator {
	return &Platform{
		rect:  rect,
		color: color,
	}
}

func (p *Platform) Generate(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(p.rect.Min, p.rect.Max)
	imd.Rectangle(0)
}
