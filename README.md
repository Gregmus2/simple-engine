Engine to create small simulations.

It is providing 
* Graphics (opengl-v4.6 + glfw-v3.3) https://github.com/go-gl/gl
* Physics (box2d) https://github.com/ByteArena/box2d

![](example.gif)

### Requirements

```build-essential libgl1-mesa-dev xorg-dev```

### Usage

* With your buildContainer method you can provide additional structures into your application
  (actually into dig service container (go.uber.org/dig))
* You should implement Scene interface.
You can use `scenes.Base` as embedded structure
* Also, you can implement `common.Init` interface to modify OpenGL settings with `Decorate`
uber method instead of `Provide`.

You can check https://github.com/Gregmus2/simple-engine-example as an example. 