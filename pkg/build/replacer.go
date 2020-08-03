package build

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	g "github.com/kesslerdev/gitgen/pkg/generator"
	"github.com/rs/zerolog/log"
)

// ReplacerCaseChanger is used to change input case
type ReplacerCaseChanger func(s string) string

type ReplacerInfos struct {
	Case string
	Var  string
}

var cases = []string{
	"exact", "upper", "lower",
	"camel", "camel_up",
	"underscore", "constant",
	"kebab", "kebab_uppercase",
}

func ApplyReplacer(r *g.BuildReplacerSpec, content []byte, f func(*ReplacerInfos) string) []byte {
	if r.Cases {
		// foreach cases
		for _, c := range cases {
			s := r.Find
			switch c {
			case "upper":
				s = strings.ToUpper(s)
			case "lower":
				s = strings.ToLower(s)
			case "camel":
				s = strcase.ToLowerCamel(s)
			case "camel_up":
				s = strcase.ToCamel(s)
			case "underscore":
				s = strcase.ToSnake(s)
			case "constant":
				s = strcase.ToScreamingSnake(s)
			case "kebab":
				s = strcase.ToKebab(s)
			case "kebab_uppercase":
				s = strcase.ToScreamingKebab(s)
			}
			content = applyTransform(&g.BuildReplacerSpec{
				As:   r.As,
				When: r.When,
				Find: s,
			}, f(&ReplacerInfos{
				Var:  r.As,
				Case: c,
			}), content)
		}
	} else {
		content = applyTransform(r, f(&ReplacerInfos{
			Var:  r.As,
			Case: "exact",
		}), content)
	}

	return content
}

func applyTransform(r *g.BuildReplacerSpec, replace string, into []byte) []byte {
	re := regexp.MustCompile(regexp.QuoteMeta(r.Find))

	if r.When != "" {
		rest := strings.Replace(r.When, r.Find, "", -1)
		re = regexp.MustCompile(regexp.QuoteMeta(r.When))
		replace = fmt.Sprint(replace, rest)
	}
	log.Trace().Msgf("Apply replacer using regex /%s/g,\nreplace with %s", re.String(), replace)
	return re.ReplaceAll(into, []byte(replace))
}
