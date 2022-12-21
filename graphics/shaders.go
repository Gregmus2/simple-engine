package graphics

const defaultVertexShaderSource string = `
    #version 420
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

const defaultFragmentShaderTemplate string = `
    #version 420
    out vec4 frag_colour;
	uniform vec3 color;
    void main() {
        frag_colour = vec4(color, 1.0);
    }
` + "\x00"

const textVertexShaderSource string = `
	#version 420
	layout (location = 0) in vec4 vertex; // <vec2 pos, vec2 tex>
	out vec2 TexCoords;

	//uniform mat4 projection;

	void main()
	{
		//gl_Position = projection * vec4(vertex.xy, 0.0, 1.0);
		gl_Position = vec4(vertex.xy, 0.0, 1.0);
		TexCoords = vertex.zw;
	}
` + "\x00"

const textFragmentShaderTemplate string = `
	#version 420
	in vec2 TexCoords;
	out vec4 frag_colour;
	
	layout (binding = 0) uniform sampler2D text; // bitmap image
	uniform vec3 color;
	
	void main()
	{    
		vec4 sampled = vec4(1.0, 1.0, 1.0, texture(text, TexCoords).r);
		frag_colour = vec4(color, 1.0) * sampled;
	} 
` + "\x00"
