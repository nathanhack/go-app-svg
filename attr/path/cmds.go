package path

import (
	svg "github.com/nathanhack/go-app-svg"
	"github.com/nathanhack/go-app-svg/internal"
)

type Cmd string

// M is MoveTo where x,y are coordinates
func M(x, y svg.Number) Cmd {
	return Cmd("M " + internal.Stringify(x, " ") + internal.Stringify(y, " "))
}

// MBy is MoveTo where dx,dy are change in coordinates
func MBy(dx, dy svg.Number) Cmd {
	return Cmd("m " + internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// L is LineTo where x,y are coordinates
func L(x, y svg.Number) Cmd {
	return Cmd("L " + internal.Stringify(x, " ") + internal.Stringify(y, " "))

}

// LBy is LineTo where dx,dy are change in coordinates
func LBy(dx, dy svg.Number) Cmd {
	return Cmd("l " + internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// H is a horizontal LineTo
func H(x svg.Number) Cmd {
	return Cmd("H " + internal.Stringify(x, " "))

}

// HBy is a horizontal LineTo where dx is the change in coordinate
func HBy(dx svg.Number) Cmd {
	return Cmd("h " + internal.Stringify(dx, " "))

}

// V is a vertical LineTo
func V(y svg.Number) Cmd {
	return Cmd("V " + internal.Stringify(y, " "))

}

// VBy is LineTo where dx,dy are change in coordinates
func VBy(dy svg.Number) Cmd {
	return Cmd("v " + internal.Stringify(dy, " "))

}

// C is cubic bezier curve
func C(x1, y1, x2, y2, x, y svg.Number) Cmd {
	return Cmd("C " + internal.Stringify(x1, " ") + internal.Stringify(y1, " ") +
		internal.Stringify(x2, " ") + internal.Stringify(y2, " ") +
		internal.Stringify(x, " ") + internal.Stringify(y, " "))

}

// CBy is cubic bezier curve
func CBy(dx1, dy1, dx2, dy2, dx, dy svg.Number) Cmd {
	return Cmd("c " + internal.Stringify(dx1, " ") + internal.Stringify(dy1, " ") +
		internal.Stringify(dx2, " ") + internal.Stringify(dy2, " ") +
		internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// S is a smooth cubic bezier curve
func S(x2, y2, x, y svg.Number) Cmd {
	return Cmd("S " + internal.Stringify(x2, " ") + internal.Stringify(y2, " ") +
		internal.Stringify(x, " ") + internal.Stringify(y, " "))

}

// SBy is a smooth cubic bezier curve
func SBy(dx2, dy2, dx, dy svg.Number) Cmd {
	return Cmd("s " + internal.Stringify(dx2, " ") + internal.Stringify(dy2, " ") +
		internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// Q is a quadratic bezier curve
func Q(x1, y1, x, y svg.Number) Cmd {
	return Cmd("Q " + internal.Stringify(x1, " ") + internal.Stringify(y1, " ") +
		internal.Stringify(x, " ") + internal.Stringify(y, " "))

}

// QBy is a quadratic bezier curve
func QBy(dx1, dy1, dx, dy svg.Number) Cmd {
	return Cmd("q " + internal.Stringify(dx1, " ") + internal.Stringify(dy1, " ") +
		internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// T is a smooth quadratic bezier curve
func T(x, y svg.Number) Cmd {
	return Cmd("T " + internal.Stringify(x, " ") + internal.Stringify(y, " "))

}

// TBy is a smooth quadratic bezier curve
func TBy(dx1, dy1, dx, dy svg.Number) Cmd {
	return Cmd("t " + internal.Stringify(dx1, " ") + internal.Stringify(dy1, " ") +
		internal.Stringify(dx, " ") + internal.Stringify(dy, " "))

}

// A an arc curve
// largeArcFlag: large(1), or small(0)
// sweepFlag: clockwise(1), or anticlockwise(0)
func A(rx, ry, angle, largeArcFlag, sweepFlag, x, y svg.Number) Cmd {
	return Cmd("A " +
		internal.Stringify(rx, " ") +
		internal.Stringify(ry, " ") +
		internal.Stringify(angle, " ") +
		internal.Stringify(largeArcFlag, " ") +
		internal.Stringify(sweepFlag, " ") +
		internal.Stringify(x, " ") +
		internal.Stringify(y, " "))
}

// ABy is an arc curve
// largeArcFlag: large(1), or small(0)
// sweepFlag: clockwise(1), or anticlockwise(0)
func ABy(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy svg.Number) Cmd {
	return Cmd("a " +
		internal.Stringify(rx, " ") +
		internal.Stringify(ry, " ") +
		internal.Stringify(angle, " ") +
		internal.Stringify(largeArcFlag, " ") +
		internal.Stringify(sweepFlag, " ") +
		internal.Stringify(dx, " ") +
		internal.Stringify(dy, " "))
}

// Z closes the current subpaht by connecting the last point of the path with it's initial point via a line
func Z() Cmd {
	return "Z "
}

// M is MoveTo where x,y are coordinates
func (c Cmd) M(x, y svg.Number) Cmd {
	return c + M(x, y)
}

// MBy is MoveTo where dx,dy are change in coordinates
func (c Cmd) MBy(dx, dy svg.Number) Cmd {
	return c + MBy(dx, dy)
}

// L is LineTo where x,y are coordinates
func (c Cmd) L(x, y svg.Number) Cmd {
	return c + L(x, y)
}

// LBy is LineTo where dx,dy are change in coordinates
func (c Cmd) LBy(dx, dy svg.Number) Cmd {
	return c + LBy(dx, dy)
}

// H is a horizontal LineTo
func (c Cmd) H(x svg.Number) Cmd {
	return c + H(x)
}

// HBy is a horizontal LineTo where dx is the change in coordinate
func (c Cmd) HBy(dx svg.Number) Cmd {
	return c + HBy(dx)
}

// V is a vertical LineTo
func (c Cmd) V(y svg.Number) Cmd {
	return c + V(y)
}

// VBy is LineTo where dx,dy are change in coordinates
func (c Cmd) VBy(dy svg.Number) Cmd {
	return c + VBy(dy)
}

// C is cubic bezier curve
func (c Cmd) C(x1, y1, x2, y2, x, y svg.Number) Cmd {
	return c + C(x1, y1, x2, y2, x, y)

}

// CBy is cubic bezier curve
func (c Cmd) CBy(dx1, dy1, dx2, dy2, dx, dy svg.Number) Cmd {
	return c + CBy(dx1, dy1, dx2, dy2, dx, dy)
}

// S is a smooth cubic bezier curve
func (c Cmd) S(x2, y2, x, y svg.Number) Cmd {
	return c + S(x2, y2, x, y)
}

// SBy is a smooth cubic bezier curve
func (c Cmd) SBy(dx2, dy2, dx, dy svg.Number) Cmd {
	return c + SBy(dx2, dy2, dx, dy)
}

// Q is a quadratic bezier curve
func (c Cmd) Q(x1, y1, x, y svg.Number) Cmd {
	return c + Q(x1, y1, x, y)
}

// QBy is a quadratic bezier curve
func (c Cmd) QBy(dx1, dy1, dx, dy svg.Number) Cmd {
	return c + QBy(dx1, dy1, dx, dy)
}

// T is a smooth quadratic bezier curve
func (c Cmd) T(x, y svg.Number) Cmd {
	return c + T(x, y)
}

// TBy is a smooth quadratic bezier curve
func (c Cmd) TBy(dx1, dy1, dx, dy svg.Number) Cmd {
	return c + TBy(dx1, dy1, dx, dy)
}

// A an arc curve
// largeArcFlag: large(1), or small(0)
// sweepFlag: clockwise(1), or anticlockwise(0)
func (c Cmd) A(rx, ry, angle, largeArcFlag, sweepFlag, x, y svg.Number) Cmd {
	return c + A(rx, ry, angle, largeArcFlag, sweepFlag, x, y)
}

// ABy is an arc curve
// largeArcFlag: large(1), or small(0)
// sweepFlag: clockwise(1), or anticlockwise(0)
func (c Cmd) ABy(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy svg.Number) Cmd {
	return c + ABy(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy)
}

// Z closes the current subpaht by connecting the last point of the path with it's initial point via a line
func (c Cmd) Z() Cmd {
	return c + Z()
}

func MoveTo(x, y svg.Number) Cmd {
	return M(x, y)
}

func MoveToRel(dx, dy svg.Number) Cmd {
	return MBy(dx, dy)
}

func LineTo(x, y svg.Number) Cmd {
	return L(x, y)
}

func LineToRel(dx, dy svg.Number) Cmd {
	return LBy(dx, dy)
}

func HorizontalLine(x svg.Number) Cmd {
	return H(x)
}

func HorizontalLineRel(dx svg.Number) Cmd {
	return HBy(dx)
}

func VerticalLine(y svg.Number) Cmd {
	return V(y)
}

func VerticalLineRel(dy svg.Number) Cmd {
	return VBy(dy)
}

func CubicCurve(x1, y1, x2, y2, x, y svg.Number) Cmd {
	return C(x1, y1, x2, y2, x, y)
}

func CubicCurveRel(dx1, dy1, dx2, dy2, dx, dy svg.Number) Cmd {
	return CBy(dx1, dy1, dx2, dy2, dx, dy)
}

func SmoothCubicCurve(x2, y2, x, y svg.Number) Cmd {
	return S(x2, y2, x, y)
}

func SmoothCubicCurveRel(dx2, dy2, dx, dy svg.Number) Cmd {
	return SBy(dx2, dy2, dx, dy)
}

func QuadraticCurve(x1, y1, x, y svg.Number) Cmd {
	return Q(x1, y1, x, y)
}

func QuadraticCurveRel(dx1, dy1, dx, dy svg.Number) Cmd {
	return QBy(dx1, dy1, dx, dy)
}

func SmoothQuadraticCurve(x, y svg.Number) Cmd {
	return T(x, y)
}

func SmoothQuadraticCurveRel(dx1, dy1, dx, dy svg.Number) Cmd {
	return TBy(dx1, dy1, dx, dy)
}

func Arc(rx, ry, angle, largeArcFlag, sweepFlag, x, y svg.Number) Cmd {
	return A(rx, ry, angle, largeArcFlag, sweepFlag, x, y)
}

func ArcRel(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy svg.Number) Cmd {
	return ABy(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy)
}

func ClosePath() Cmd {
	return Z()
}

func (c Cmd) MoveTo(x, y svg.Number) Cmd {
	return c + M(x, y)
}

func (c Cmd) MoveToRel(dx, dy svg.Number) Cmd {
	return c + MBy(dx, dy)
}

func (c Cmd) LineTo(x, y svg.Number) Cmd {
	return c + L(x, y)
}

func (c Cmd) LineToRel(dx, dy svg.Number) Cmd {
	return c + LBy(dx, dy)
}

func (c Cmd) HorizontalLine(x svg.Number) Cmd {
	return c + H(x)
}

func (c Cmd) HorizontalLineRel(dx svg.Number) Cmd {
	return c + HBy(dx)
}

func (c Cmd) VerticalLine(y svg.Number) Cmd {
	return c + V(y)
}

func (c Cmd) VerticalLineRel(dy svg.Number) Cmd {
	return c + VBy(dy)
}

func (c Cmd) CubicCurve(x1, y1, x2, y2, x, y svg.Number) Cmd {
	return c + C(x1, y1, x2, y2, x, y)
}

func (c Cmd) CubicCurveRel(dx1, dy1, dx2, dy2, dx, dy svg.Number) Cmd {
	return c + CBy(dx1, dy1, dx2, dy2, dx, dy)
}

func (c Cmd) SmoothCubicCurve(x2, y2, x, y svg.Number) Cmd {
	return c + S(x2, y2, x, y)
}

func (c Cmd) SmoothCubicCurveRel(dx2, dy2, dx, dy svg.Number) Cmd {
	return c + SBy(dx2, dy2, dx, dy)
}

func (c Cmd) QuadraticCurve(x1, y1, x, y svg.Number) Cmd {
	return c + Q(x1, y1, x, y)
}

func (c Cmd) QuadraticCurveRel(dx1, dy1, dx, dy svg.Number) Cmd {
	return c + QBy(dx1, dy1, dx, dy)
}

func (c Cmd) SmoothQuadraticCurve(x, y svg.Number) Cmd {
	return c + T(x, y)
}

func (c Cmd) SmoothQuadraticCurveRel(dx1, dy1, dx, dy svg.Number) Cmd {
	return c + TBy(dx1, dy1, dx, dy)
}

func (c Cmd) Arc(rx, ry, angle, largeArcFlag, sweepFlag, x, y svg.Number) Cmd {
	return c + A(rx, ry, angle, largeArcFlag, sweepFlag, x, y)
}

func (c Cmd) ArcRel(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy svg.Number) Cmd {
	return c + ABy(rx, ry, angle, largeArcFlag, sweepFlag, dx, dy)
}

func (c Cmd) ClosePath() Cmd {
	return c + Z()
}
