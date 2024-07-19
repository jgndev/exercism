package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if v, ok := cardValues[card]; ok {
		return v
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	first := ParseCard(card1)
	second := ParseCard(card2)
	dealer := ParseCard(dealerCard)

	hand := first + second

	switch {
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		return "P"

	// If you have 21 and the dealer does not have an ace, or a card value
	// of 10 (face or 10) then you win.
	case hand == 21 && dealer < 10:
		return "W"

	// If you have 21 and the dealer has an ace, or a 10 (face card or 10)
	// then stand and wait for the reveal of the other card.
	case hand == 21 && dealer >= 10:
		return "S"

	// If your card values are in the range of 17 to 20 you should stand.
	case hand >= 17 && hand <= 20:
		return "S"

	// If your card values are in the range of 12 to 16 you should always stand
	// unless the dealer has a 7 or higher, in which case you should hit.
	case hand >= 12 && hand <= 16 && dealer >= 7:
		return "H"

	// If your card values are in the range of 12 to 16 you should always stand
	// unless the dealer has a 7 or higher, in which case you should hit.
	case hand >= 12 && hand <= 16 && dealer <= 6:
		return "S"

	// If your cards sum up to 11 or lower you should always hit.
	default:
		return "H"

	}

}
