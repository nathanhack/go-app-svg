package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	svg "github.com/nathanhack/go-app-svg"
	"github.com/nathanhack/go-app-svg/attr"
	"github.com/nathanhack/go-app-svg/attr/transforms"
)

type pie struct {
	app.Compo
}

func (h *pie) Render() app.UI {
	return app.Div().Body(
		svg.Svg(
			attr.Width(100),
			attr.Height(100),
			attr.ViewBox(0, 0, 20, 20),
			attr.Fill("black"),
			svg.Circle(
				attr.Cx(10),
				attr.Cy(10),
				attr.R(10),
			),
			svg.Circle(
				attr.Cx(10),
				attr.Cy(10),
				attr.R(5),
				attr.Stroke("#FF6347"),
				attr.StrokeWidth(10),
				attr.StrokeDasharray(35*31.4/100, 31.4),
				attr.Transform(transforms.Rotate(-90, 10, 10)),
			),
		),
	)
}

func main() {
	// Components routing:
	app.Route("/", &pie{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Pie",
		Description: "Example",
	})

	port := 8000
	fmt.Printf("http://localhost:%v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
