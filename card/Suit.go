package card

import (
	"github.com/RoboSkyrat/pickup/util"
	"strings"
)

type Suit int

const (
	Clubs Suit	= 100
	Diamonds	= 200
	Hearts		= 300
	Spades		= 400
)

func GetSuit(face string) Suit {
	face = strings.ToLower(face)
	return (Suit(util.IndexOf(face, []string{
		"clubs",
		"diamonds",
		"hearts",
		"spades",
	}))+1)*100
}