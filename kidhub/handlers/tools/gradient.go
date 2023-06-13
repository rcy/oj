package tools

import (
	"fmt"
	"html/template"
	"strings"
)

type Stop struct {
	Color   string
	Percent int
}

type Gradient struct {
	Type    string
	Repeat  bool
	Stops   []Stop
	Degrees int
}

func (g Gradient) Render() template.CSS {
	return g.render(g.Type, g.Repeat, g.Degrees, g.Stops)
}

func (g Gradient) RenderBar() template.CSS {
	return g.render("linear", false, 90, g.Stops)
}

func (g Gradient) render(gradientType string, repeating bool, deg int, stops []Stop) template.CSS {
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
