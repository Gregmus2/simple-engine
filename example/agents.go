package main

import (
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/scenes"
	"github.com/gregmus2/nnga"
	"github.com/patrikeh/go-deep"
	"math/rand"
	"time"
)

type Agents struct {
	scenes.Base
	factory *ObjectFactory
	con     *graphics.PercentToPosConverter
}

func NewAgents(base scenes.Base, f *ObjectFactory, con *graphics.PercentToPosConverter) *Agents {
	return &Agents{
		Base:    base,
		factory: f,
		con:     con,
	}
}

func (d *Agents) Init() {
	rand.Seed(time.Now().UTC().UnixNano())

	ga := nnga.NewGA(1000, &deep.Config{
		/* Input dimensionality */
		Inputs: 2,
		/* Two hidden layers consisting of two neurons each, and a single output */
		Layout: []int{2, 2, 1},
		/* Activation functions: Sigmoid, Tanh, ReLU, Linear */
		Activation: deep.ActivationSigmoid,
		/* Determines output layer activation & loss function:
		ModeRegression: linear outputs with MSE loss
		ModeMultiClass: softmax output with Cross Entropy loss
		ModeMultiLabel: sigmoid output with Cross Entropy loss
		ModeBinary: sigmoid output with binary CE loss */
		Mode: deep.ModeBinary,
		/* Weight initializers: {deep.NewNormal(μ, σ), deep.NewUniform(μ, σ)} */
		Weight: deep.NewNormal(1.0, 0.0),
		/* Apply bias */
		Bias: true,
	}, &nnga.Coefficients{
		Scale:                   1,
		Selection:               0.2,
		MutationClassic:         0.1,
		MutationGrowth:          2,
		MutationGenesMaxPercent: 0.2,
		MutationOffset:          0.1,
	})

	for i := 0; i < 50; i++ {
		food := d.factory.NewFood(float64(rand.Intn(d.Cfg.Window.W)), float64(rand.Intn(d.Cfg.Window.H)))
		d.DrawObjects = append(d.DrawObjects, food)
	}

	agent := d.factory.NewAgent(d.con.X(50), d.con.Y(50))
	d.DrawObjects = append(d.DrawObjects, agent)
}
