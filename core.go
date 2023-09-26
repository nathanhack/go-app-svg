package svg

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/nathanhack/go-app-svg/internal"
)

var Namespace = "http://www.w3.org/2000/svg"

type Number interface{}

type Length interface{}

type LengthOrPercent interface{}

type NumberOrPercent interface{}

func Percent(number Number) interface{} {
	return internal.Stringify(number, "%")
}

func If(value bool, trueStringAttrOrElem any, falseStringAttrOrElem any) any {
	if value {
		return trueStringAttrOrElem
	}
	return falseStringAttrOrElem
}

func IfSlice(value bool, trueStringAttrOrElem []any, falseStringAttrOrElem []any) []any {
	if value {
		return trueStringAttrOrElem
	}
	return falseStringAttrOrElem
}

type Attribute interface {
	Name() string
	Value() any
}

// Element is for creating generic elements not currently defined in the package
func Element(tag string, stringAttrsOrElements ...any) app.HTMLElem {
	elements := make([]app.UI, 0)

	ret := app.Elem(tag)
	ret.XMLNS(Namespace)

	var processType func(value any)

	processType = func(value any) {
		switch v := value.(type) {
		case []Attribute:
			for _, a := range v {
				processType(a)
			}
		case []any:
			for _, a := range v {
				processType(a)
			}
		case app.HTMLElem:
			elements = append(elements, v)
		case app.UI:
			elements = append(elements, v)

		case Attribute:
			ret.Attr(v.Name(), v.Value())
		case string:
			elements = append(elements, app.Raw(v))
		default:
			panic(fmt.Sprintf("Unsupported type: %T", value))
		}
	}

	for _, sae := range stringAttrsOrElements {
		processType(sae)
	}
	ret.Body(elements...)
	return ret
}

func attrsOnly(tag string, attrs ...Attribute) app.HTMLElem {
	ret := app.Elem(tag)
	ret.XMLNS(Namespace)

	for _, x := range attrs {
		tmp := x.(Attribute)
		ret.Attr(tmp.Name(), tmp.Value())
	}
	return ret
}

func Svg(stringAttrsOrElements ...any) app.HTMLElem {
	return Element("svg", stringAttrsOrElements...)
}

///////////////////////////////////////////////////

func A(attrsOrElements ...any) app.HTMLElem {
	return Element("a", attrsOrElements...)
}

func Animate(attrsOrElements ...any) app.HTMLElem {
	return Element("animate", attrsOrElements...)
}

func AnimateMotion(attrsOrElements ...any) app.HTMLElem {
	return Element("animateMotion", attrsOrElements...)
}
func Circle(attrsOrElements ...any) app.HTMLElem {
	return Element("circle", attrsOrElements...)
}

func ClipPath(attrsOrElements ...any) app.HTMLElem {
	return Element("clipPath", attrsOrElements...)
}

func Defs(elements ...app.UI) app.HTMLElem {
	defs := app.Elem("defs")
	defs.XMLNS(Namespace)
	defs.Body(elements...)
	return defs
}

// Desc provides an accessible, long-text description of any SVG container element or graphics element.
func Desc(desc string) app.HTMLElem {
	d := Element("desc")
	d.Body(app.Text(desc))
	return d
}

// Discard allows authors to specify the time at which particular elements are to be discarded,
// thereby reducing the resources required by an SVG user agent.
func Discard(attrs ...Attribute) app.HTMLElem {
	return attrsOnly("discard", attrs...)
}

func Ellipse(attrsOrElements ...any) app.HTMLElem {
	return Element("ellipse", attrsOrElements...)
}

func FeBlend(attrsOrElements ...any) app.HTMLElem {
	return Element("feBlend", attrsOrElements...)
}

func FeColorMatrix(attrsOrElements ...any) app.HTMLElem {
	return Element("feColorMatrix", attrsOrElements...)
}

func FeComponentTransfer(attrsOrElements ...any) app.HTMLElem {
	return Element("feComponentTransfer", attrsOrElements...)
}

func FeComposite(attrsOrElements ...any) app.HTMLElem {
	return Element("feComposite", attrsOrElements...)
}

func FeConvolveMatrix(attrsOrElements ...any) app.HTMLElem {
	return Element("feConvolveMatrix", attrsOrElements...)
}

func FeDiffuseLighting(attrsOrElements ...any) app.HTMLElem {
	return Element("feDiffuseLighting", attrsOrElements...)
}

func FeDisplacementMap(attrsOrElements ...any) app.HTMLElem {
	return Element("feDisplacementMap", attrsOrElements...)
}

func FeDropShadow(attrsOrElements ...any) app.HTMLElem {
	return Element("feDropShadow", attrsOrElements...)
}

func FeDistantLight(attrsOrElements ...any) app.HTMLElem {
	return Element("feDistantLight", attrsOrElements...)
}

func FeFlood(attrsOrElements ...any) app.HTMLElem {
	return Element("feFlood", attrsOrElements...)
}

func FeFuncA(attrsOrElements ...any) app.HTMLElem {
	return Element("feFuncA", attrsOrElements...)
}

func FeFuncB(attrsOrElements ...any) app.HTMLElem {
	return Element("feFuncB", attrsOrElements...)
}

func FeFuncG(attrsOrElements ...any) app.HTMLElem {
	return Element("feFuncG", attrsOrElements...)
}

func FeFuncR(attrsOrElements ...any) app.HTMLElem {
	return Element("feFuncR", attrsOrElements...)
}

func FeGaussianBlur(attrsOrElements ...any) app.HTMLElem {
	return Element("feGaussianBlur", attrsOrElements...)
}

func FeImage(attrsOrElements ...any) app.HTMLElem {
	return Element("feImage", attrsOrElements...)
}

func FeMerge(attrsOrElements ...any) app.HTMLElem {
	return Element("feMerge", attrsOrElements...)
}

func FeMergeNode(attrsOrElements ...any) app.HTMLElem {
	return Element("feMergeNode", attrsOrElements...)
}

func FeOffset(attrsOrElements ...any) app.HTMLElem {
	return Element("feOffset", attrsOrElements...)
}

func FePointLight(attrsOrElements ...any) app.HTMLElem {
	return Element("fePointLight", attrsOrElements...)
}

func FeSpecularLighting(attrsOrElements ...any) app.HTMLElem {
	return Element("feSpecularLighting", attrsOrElements...)
}

func FeSpotLight(attrsOrElements ...any) app.HTMLElem {
	return Element("feSpotLight", attrsOrElements...)
}

func FeTile(attrsOrElements ...any) app.HTMLElem {
	return Element("feTile", attrsOrElements...)
}

func FeTurbulence(attrsOrElements ...any) app.HTMLElem {
	return Element("feTurbulence", attrsOrElements...)
}

func Filter(attrsOrElements ...any) app.HTMLElem {
	return Element("filter", attrsOrElements...)
}

func ForeignObject(attrsOrElements ...any) app.HTMLElem {
	return Element("foreignObject", attrsOrElements...)
}

func G(attrsOrElements ...any) app.HTMLElem {
	return Element("g", attrsOrElements...)
}

func Image(attrsOrElements ...any) app.HTMLElem {
	return Element("image", attrsOrElements...)
}

func Line(attrsOrElements ...any) app.HTMLElem {
	return Element("line", attrsOrElements...)
}

func LinearGradient(attrsOrElements ...any) app.HTMLElem {
	return Element("linearGradient", attrsOrElements...)
}

func Marker(attrsOrElements ...any) app.HTMLElem {
	return Element("marker", attrsOrElements...)
}

func Mask(attrsOrElements ...any) app.HTMLElem {
	return Element("mask", attrsOrElements...)
}

func Metadata(data any) app.HTMLElem {
	return Element("metadata", data)
}
func MPath(attrsOrElements ...any) app.HTMLElem {
	return Element("mpath", attrsOrElements...)
}

func Path(attrsOrElements ...any) app.HTMLElem {
	return Element("path", attrsOrElements...)
}

func Pattern(attrsOrElements ...any) app.HTMLElem {
	return Element("pattern", attrsOrElements...)
}

func Polyline(attrsOrElements ...any) app.HTMLElem {
	return Element("polyline", attrsOrElements...)
}

func Polygon(attrsOrElements ...any) app.HTMLElem {
	return Element("polygon", attrsOrElements...)
}

func RadialGradient(attrsOrElements ...any) app.HTMLElem {
	return Element("radialGradient", attrsOrElements...)
}

func Rect(attrsOrElements ...any) app.HTMLElem {
	return Element("rect", attrsOrElements...)
}

func Set(attrs ...Attribute) app.HTMLElem {
	return attrsOnly("set", attrs...)
}

func Stop(attrsOrElements ...any) app.HTMLElem {
	return Element("stop", attrsOrElements...)
}

func Switch(attrsOrElements ...any) app.HTMLElem {
	return Element("switch", attrsOrElements...)
}

func Symbol(attrsOrElements ...any) app.HTMLElem {
	return Element("symbol", attrsOrElements...)
}

func Text(stringAttrsOrElements ...any) app.HTMLElem {
	return Element("text", stringAttrsOrElements...)
}

func TextPath(stringAttrsOrElements ...any) app.HTMLElem {
	return Element("textPath", stringAttrsOrElements...)
}

func Title(value string, attrs ...Attribute) app.HTMLElem {
	a := attrsOnly("title", attrs...)
	a.Body(app.Raw(value))
	return a
}

func Tspan(value string, stringAttrsOrElements ...any) app.HTMLElem {
	return Element("tspan", stringAttrsOrElements...)
}

func Use(attrs ...Attribute) app.HTMLElem {
	return attrsOnly("use", attrs...)
}

func View(attrs ...Attribute) app.HTMLElem {
	return attrsOnly("view", attrs...)
}
