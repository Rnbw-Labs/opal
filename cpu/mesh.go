package cpu

import (
	"github.com/rnbw-labs/opal/gmath"
)

type Mesh struct {
	Vertices  []Vertex
	Triangles []Triangle
}

func (m *Mesh) Merge(o Mesh) {
	vertexOffset := uint32(len(m.Vertices))
	m.Vertices = append(m.Vertices, o.Vertices...)
	newTris := make([]Triangle, len(o.Triangles))
	for i, t := range o.Triangles {
		var newIndices [3]uint32
		for k, v := range t.Indices {
			newIndices[k] = v + vertexOffset
		}
		newTris[i] = Triangle{Indices: newIndices}
	}
	m.Triangles = append(m.Triangles, newTris...)
}

type Triangle struct {
	Indices [3]uint32
}

type Vertex struct {
	Position gmath.Vector3
	Normal   gmath.Vector3
	UV       gmath.Vector2
}

func (v Vertex) Unpack() []float32 {
	return []float32{
		v.Position.X, v.Position.Y, v.Position.Z,
		v.Normal.X, v.Normal.Y, v.Normal.Z,
		v.UV.X, v.UV.Y,
	}
}

func CubeMesh(size float32) *Mesh {
	h := size / 2

	// 6 faces, each with its own 4 vertices (for correct normals + UVs)
	vertices := []Vertex{
		// Front face (+Z)
		{Position: gmath.Vector3{-h, -h, h}, Normal: gmath.Vector3{0, 0, 1}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{h, -h, h}, Normal: gmath.Vector3{0, 0, 1}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{h, h, h}, Normal: gmath.Vector3{0, 0, 1}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{-h, h, h}, Normal: gmath.Vector3{0, 0, 1}, UV: gmath.Vector2{0, 1}},

		// Back face (-Z)
		{Position: gmath.Vector3{h, -h, -h}, Normal: gmath.Vector3{0, 0, -1}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{-h, -h, -h}, Normal: gmath.Vector3{0, 0, -1}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{-h, h, -h}, Normal: gmath.Vector3{0, 0, -1}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{h, h, -h}, Normal: gmath.Vector3{0, 0, -1}, UV: gmath.Vector2{0, 1}},

		// Left face (-X)
		{Position: gmath.Vector3{-h, -h, -h}, Normal: gmath.Vector3{-1, 0, 0}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{-h, -h, h}, Normal: gmath.Vector3{-1, 0, 0}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{-h, h, h}, Normal: gmath.Vector3{-1, 0, 0}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{-h, h, -h}, Normal: gmath.Vector3{-1, 0, 0}, UV: gmath.Vector2{0, 1}},

		// Right face (+X)
		{Position: gmath.Vector3{h, -h, h}, Normal: gmath.Vector3{1, 0, 0}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{h, -h, -h}, Normal: gmath.Vector3{1, 0, 0}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{h, h, -h}, Normal: gmath.Vector3{1, 0, 0}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{h, h, h}, Normal: gmath.Vector3{1, 0, 0}, UV: gmath.Vector2{0, 1}},

		// Top face (+Y)
		{Position: gmath.Vector3{-h, h, h}, Normal: gmath.Vector3{0, 1, 0}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{h, h, h}, Normal: gmath.Vector3{0, 1, 0}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{h, h, -h}, Normal: gmath.Vector3{0, 1, 0}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{-h, h, -h}, Normal: gmath.Vector3{0, 1, 0}, UV: gmath.Vector2{0, 1}},

		// Bottom face (-Y)
		{Position: gmath.Vector3{-h, -h, -h}, Normal: gmath.Vector3{0, -1, 0}, UV: gmath.Vector2{0, 0}},
		{Position: gmath.Vector3{h, -h, -h}, Normal: gmath.Vector3{0, -1, 0}, UV: gmath.Vector2{1, 0}},
		{Position: gmath.Vector3{h, -h, h}, Normal: gmath.Vector3{0, -1, 0}, UV: gmath.Vector2{1, 1}},
		{Position: gmath.Vector3{-h, -h, h}, Normal: gmath.Vector3{0, -1, 0}, UV: gmath.Vector2{0, 1}},
	}

	// 6 faces × 2 triangles each
	triangles := []Triangle{
		// Front
		{Indices: [3]uint32{0, 1, 2}},
		{Indices: [3]uint32{0, 2, 3}},

		// Back
		{Indices: [3]uint32{4, 5, 6}},
		{Indices: [3]uint32{4, 6, 7}},

		// Left
		{Indices: [3]uint32{8, 9, 10}},
		{Indices: [3]uint32{8, 10, 11}},

		// Right
		{Indices: [3]uint32{12, 13, 14}},
		{Indices: [3]uint32{12, 14, 15}},

		// Top
		{Indices: [3]uint32{16, 17, 18}},
		{Indices: [3]uint32{16, 18, 19}},

		// Bottom
		{Indices: [3]uint32{20, 21, 22}},
		{Indices: [3]uint32{20, 22, 23}},
	}

	return &Mesh{
		Vertices:  vertices,
		Triangles: triangles,
	}
}
