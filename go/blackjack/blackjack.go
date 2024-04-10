package blackjack

var cardValues = map[string]int{
    "ace":   11,
    "king":  10,
    "queen": 10,
    "jack":  10,
    "ten":   10,
    "nine":  9,
    "eight": 8,
    "seven": 7,
    "six":   6,
    "five":  5,
    "four":  4,
    "three": 3,
    "two":   2,
}

const (
    split = "P" // Player chooses to split
    stand = "S" // Player chooses to stand
    hit   = "H" // Player chooses to hit
    win   = "W" // Player wins
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    c1 := cardValues[card1]
    c2 := cardValues[card2]
    hand := cardValues[card1] + cardValues[card2]

    dc := cardValues[dealerCard]

    switch {

    case c1 == 11 && c2 == 11:
        return split

    case hand == 21 && dc < 10:
        return win

    case hand == 21 && dc >= 10:
        return stand

    case hand >= 17 && hand <= 20:
        return stand

    case hand >= 12 && hand <= 16 && dc >= 7:
        return hit

    case hand >= 12 && hand <= 16:
        return stand

    default:
        return hit
    }
}
