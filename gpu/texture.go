package gpu

import "github.com/go-gl/gl/v4.6-core/gl"

type Texture struct {
	ID uint32
}

func (t *Texture) Bind(slot uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + slot)
	gl.BindTexture(gl.TEXTURE_2D, t.ID)
}
