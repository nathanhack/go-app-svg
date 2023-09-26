# go-app-svg
Svg support for [go-app](https://github.com/maxence-charriere/go-app)

# Example

```
func (c*circle) Render() app.UI {
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
		),
	)
}
```