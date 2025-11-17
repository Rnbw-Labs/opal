package gpu

import "github.com/go-gl/gl/v4.6-core/gl"

type Mesh struct {
	VAO        uint32
	VBO        uint32
	EBO        uint32
	IndexCount uint32
}

func (m *Mesh) Draw() {
	gl.BindVertexArray(m.VAO)
	gl.DrawElements(
		gl.TRIANGLES,
		int32(m.IndexCount),
		gl.UNSIGNED_INT,
		nil,
	)
}
