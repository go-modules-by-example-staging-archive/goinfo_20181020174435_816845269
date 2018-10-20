package contributors

var all = [...]string{
	"Robert Griesemer",
	"Rob Pike",
	"Ken Thompson",
	"Russ Cox",
	"Ian Lance Taylor",
}

func Names() []string {
	res := all
	return res[:]
}
