package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	svg "github.com/nathanhack/go-app-svg"
	"github.com/nathanhack/go-app-svg/attr"
)

type circle struct {
	app.Compo

	clicked bool
}

func (c *circle) OnClick(ctx app.Context, e app.Event) {
	c.clicked = !c.clicked
	c.Render()
}

func (c *circle) Render() app.UI {
	return app.Div().Body(
		svg.Svg(
			attr.Width(300),
			attr.Height(100),
			attr.ViewBox(0, 0, 300, 100),
			attr.If(c.clicked, attr.Fill("grey"), attr.Fill("black")),
			attr.Stroke("red"),
			svg.Circle(
				attr.Cx(50),
				attr.Cy(50),
				attr.R(40),
			),
			svg.If(c.clicked, svg.Circle(
				attr.Cx(150),
				attr.Cy(50),
				attr.R(4),
				attr.Fill("purple"),
			), nil),
		),
	)
}

func main() {
	// Components routing:
	app.Route("/", &circle{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Circle",
		Description: "Example",
	})

	port := 8000
	fmt.Printf("http://localhost:%v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
