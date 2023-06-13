package gradient

import (
	"fmt"
	"html/template"
	"math/rand"
	"sort"
	"strings"
)

type stop struct {
	Color   string
	Percent int
}

type Gradient struct {
	Type      string
	Repeat    bool
	Degrees   int
	Colors    []string
	Positions []int
}

var Default = Gradient{
	Type:      "linear",
	Repeat:    false,
	Degrees:   90,
	Colors:    []string{"#ff00ff", "#ffff00", "#00ffff"},
	Positions: []int{0, 50, 100},
}

var Black = Gradient{
	Type:      "linear",
	Repeat:    false,
	Degrees:   90,
	Colors:    []string{"#000000"},
	Positions: []int{50},
}

func Random() Gradient {
	n := rand.Intn(3) + 2
	colors := make([]string, n)
	positions := make([]int, n)
	for i := 0; i < n; i += 1 {
		colors[i] = randHexColor()
		positions[i] = rand.Intn(100)
	}
	// colors[n-1] = colors[0]
	// positions[n-1] = 100
	// positions[0] = 0

	return Gradient{
		Type:      []string{"linear", "radial", "conic"}[rand.Intn(3)],
		Repeat:    []bool{false, true}[rand.Intn(2)],
		Degrees:   rand.Intn(180),
		Colors:    colors,
		Positions: positions,
	}
}

func randHexColor() string {
	return fmt.Sprintf("#%x", rand.Uint32())[0:7]
}

func (g Gradient) Stops() []stop {
	var stops []stop

	// zip colors and positions into stops
	for i, c := range g.Colors {
		p := g.Positions[i]
		stops = append(stops, stop{Color: c, Percent: p})
	}

	// sort the stops by position
	sort.Slice(stops, func(i, j int) bool {
		return stops[i].Percent < stops[j].Percent
	})

	return stops
}

// Render the gradient as a css value
func (g Gradient) Render() template.CSS {
	return g.render(g.Type, g.Repeat, g.Degrees, g.Stops())
}

// Render a gradient as a css value that can be used as a horizontal slider bar
func (g Gradient) RenderBar() template.CSS {
	return g.render("linear", false, 90, g.Stops())
}

func (g Gradient) render(gradientType string, repeating bool, deg int, stops []stop) template.CSS {
	var params []string

	if repeating {
		gradientType = "repeating-" + gradientType
	}

	switch gradientType {
	case "linear":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`linear-gradient(%ddeg, %s)`, deg, strings.Join(params, ",")))

	case "radial":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`radial-gradient(%s)`, strings.Join(params, ",")))

	case "conic":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`conic-gradient(from %ddeg, %s)`, deg, strings.Join(params, ",")))

	case "repeating-linear":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %dpx", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`repeating-linear-gradient(%ddeg, %s)`, deg, strings.Join(params, ",")))
	case "repeating-radial":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %dpx", s.Color, s.Percent))
		}
		return template.CSS(fmt.Sprintf(`repeating-radial-gradient(%s)`, strings.Join(params, ",")))
	case "repeating-conic":
		for _, s := range stops {
			params = append(params, fmt.Sprintf("%s %d%%", s.Color, s.Percent/4))
		}
		return template.CSS(fmt.Sprintf(`repeating-conic-gradient(from %ddeg, %s)`, deg, strings.Join(params, ",")))
	default:
		return template.CSS("black")
	}
}
