package gpu

var BasicVertexShaderSource = `
#version 330 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 aNormal;
layout (location = 2) in vec2 aUV;

out vec2 vUV;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main() {
    vUV = aUV;
    gl_Position = projection * view * model * vec4(aPos, 1.0);
}
`

var BasicFragmentShaderSource = `
#version 330 core

out vec4 FragColor;

void main() {
    FragColor = vec4(1.0, 1.0, 1.0, 1.0); // solid pure white
}
`
