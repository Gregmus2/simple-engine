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

const textVertexShaderSource = `
	#version 420
	in vec2 vert;
	in vec2 vertTexCoord;

	//window res
	uniform vec2 resolution;
	out vec2 fragTexCoord;

	void main() {
	   // convert the rectangle from pixels to 0.0 to 1.0
	   vec2 zeroToOne = vert / resolution;
	   // convert from 0->1 to 0->2
	   vec2 zeroToTwo = zeroToOne * 2.0;
	   // convert from 0->2 to -1->+1 (clipspace)
	   vec2 clipSpace = zeroToTwo - 1.0;
	   fragTexCoord = vertTexCoord;
	   gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
	}
` + "\x00"

const textFragmentShaderTemplate = `
	#version 420
	in vec2 fragTexCoord;
	out vec4 outputColor;
	
	uniform sampler2D tex;
	uniform vec4 textColor;
	
	void main()
	{    
		vec4 sampled = vec4(1.0, 1.0, 1.0, texture(tex, fragTexCoord).r);
		outputColor = textColor * sampled;
	}
` + "\x00"
