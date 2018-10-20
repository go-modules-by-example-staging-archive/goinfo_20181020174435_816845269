package designers

import "github.com/go-modules-by-example-staging/goinfo/contributors"

func Names() []string {
	var res []string
	for _, n := range contributors.Names() {
		switch n {
		case "Rob Pike", "Ken Thompson", "Robert Griesemer":
			res = append(res, n)
		}
	}
	return res
}
