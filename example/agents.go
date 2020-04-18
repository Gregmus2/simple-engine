package main

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/graphics"
	"github.com/Gregmus2/simple-engine/objects"
	"github.com/Gregmus2/simple-engine/scenes"
	"github.com/gregmus2/nnga"
	"github.com/patrikeh/go-deep"
	"math"
	"math/rand"
	"time"
)

type Agents struct {
	scenes.Base
	factory *ObjectFactory
	con     *graphics.PercentToPosConverter
	agents  []*Agent
	food    []*Food
	ga      *nnga.GA
}

func NewAgents(base scenes.Base, f *ObjectFactory, con *graphics.PercentToPosConverter) *Agents {
	return &Agents{
		Base:    base,
		factory: f,
		con:     con,
	}
}

func (d *Agents) Init() {
	time.Sleep(5 * time.Second)

	rand.Seed(time.Now().UTC().UnixNano())

	d.walls()

	d.ga = nnga.NewGA(100, &deep.Config{
		/* Input dimensionality */
		Inputs: 2,
		/* Two hidden layers consisting of two neurons each, and a single output */
		Layout: []int{2, 4, 2},
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
		d.food = append(d.food, food)
	}

	for _, person := range d.ga.Persons {
		agent := d.factory.NewAgent(float64(rand.Intn(d.Cfg.Window.W)), float64(rand.Intn(d.Cfg.Window.H)), person)
		d.DrawObjects = append(d.DrawObjects, agent)
		d.agents = append(d.agents, agent)
	}
}

func (d *Agents) PreUpdate() {
	for _, agent := range d.agents {
		pos := agent.circle.Body.GetPosition()
		targetPos := d.food[0].Body.GetPosition()
		min := math.Sqrt(math.Pow(pos.X-targetPos.X, 2) + math.Pow(pos.Y-targetPos.Y, 2))
		for _, piece := range d.food {
			foodPos := piece.Body.GetPosition()
			distance := math.Sqrt(math.Pow(pos.X-foodPos.X, 2) + math.Pow(pos.Y-foodPos.Y, 2))
			if distance < min {
				min = distance
				targetPos = foodPos
			}
		}

		// todo Math folding plugin
		a := min
		b := math.Sqrt(math.Pow(pos.X-float64(agent.cursor.X), 2) + math.Pow(pos.Y-float64(agent.cursor.Y), 2))
		c := math.Sqrt(math.Pow(targetPos.X-float64(agent.cursor.X), 2) + math.Pow(targetPos.Y-float64(agent.cursor.Y), 2))
		angle := math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b))
		output := agent.person.Predict([]float64{math.Cos(angle), min})
		dAngle, force := output[0], output[1]*0.1

		newAngle := agent.circle.Body.GetAngle() + dAngle
		xForce := force * math.Cos(newAngle)
		yForce := force * math.Sin(newAngle)
		agent.circle.Body.ApplyForceToCenter(box2d.MakeB2Vec2(xForce, yForce), true)
	}

	// save targetPos and distance for calc score
}

func (d *Agents) Update() {
	// calc score and evolve
}

func (d *Agents) walls() {
	b := objects.BoxModel{
		X:       d.con.X(50),
		Y:       0,
		W:       d.con.X(100),
		H:       1,
		T:       box2d.B2BodyType.B2_staticBody,
		Color:   graphics.White(),
		Density: 0,
	}
	d.DrawObjects = append(d.DrawObjects, d.factory.NewBox(b))

	b.Y = d.con.Y(100)
	d.DrawObjects = append(d.DrawObjects, d.factory.NewBox(b))

	b.X, b.Y, b.W, b.H = 0, d.con.Y(50), 1, d.con.Y(100)
	d.DrawObjects = append(d.DrawObjects, d.factory.NewBox(b))

	b.X = d.con.X(100)
	d.DrawObjects = append(d.DrawObjects, d.factory.NewBox(b))
}
