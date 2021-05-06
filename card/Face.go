package card

import (
	"github.com/RoboSkyrat/pickup/util"
	"strings"
)

type Face int

const (
	Ace	Face	= 1
	Two			= 2
	Three		= 3
	Four		= 4
	Five 		= 5
	Six			= 6
	Seven		= 7
	Eight		= 8
	Nine		= 9
	Ten			= 10
	Jack		= 11
	Queen		= 12
	King		= 13
)

func GetFace(face string) Face {
	face = strings.ToLower(face)
	return Face(util.IndexOf(face, []string{
		"ace",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"jack",
		"queen",
		"king",
	})+1)
}