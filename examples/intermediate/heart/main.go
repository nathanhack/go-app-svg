package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	svg "github.com/nathanhack/go-app-svg"
	"github.com/nathanhack/go-app-svg/attr"
	"github.com/nathanhack/go-app-svg/attr/path"
)

type heart struct {
	app.Compo
}

func (h *heart) Render() app.UI {
	return app.Div().Body(
		svg.Svg(
			attr.Width(100),
			attr.Height(100),
			attr.Stroke("red"),
			attr.Fill("grey"),
			attr.ViewBox(0, 0, 100, 100),
			svg.Path(
				attr.Fill("purple"),
				attr.Stroke("red"),
				attr.D(
					//D takes multiple paths
					// Path's can be chained together and creates a "path"
					path.M(10, 30).
						A(20, 20, 0, 0, 1, 50, 30).
						A(20, 20, 0, 0, 1, 90, 30).
						Q(90, 60, 50, 90).
						Q(10, 60, 10, 30).
						Z(), //draws a heart
					// Or two or more paths can be added together
					path.M(50, 50)+path.H(100), // draws a horizontal line to the right
					// or each path can be added separately
					// path.M(50, 50),
					path.V(100), // these two, draw a vertical line down
					//and as always literal strings work too
					"M 50 50 h -25", // another horizontal line to the left
				),
			),
		),
	)
}

func main() {
	// Components routing:
	app.Route("/", &heart{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Heart",
		Description: "Example",
	})

	port := 8000
	fmt.Printf("http://localhost:%v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
