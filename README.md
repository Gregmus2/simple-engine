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
* You should implement Scene interface to replace scenes.Demo scene.
You can use `scenes.Base` as embedded structure
* You can use `objects.ObjectFactory` structure as embedded in your factory to create new objects.
* Also you need to implement `common.Init` interface, at least with empty content

You can check https://github.com/Gregmus2/simple-engine-example as an example. 