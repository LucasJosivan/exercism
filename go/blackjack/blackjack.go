package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":    11,
		"one":    1,
		"two":    2,
		"three":  3,
		"four":   4,
		"five":   5,
		"six":    6,
		"seven":  7,
		"eight":  8,
		"nine":   9,
		"ten":    10,
		"jack":   10,
		"queen":  10,
		"king":   10,
		"jocker": 0,
	}
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	value := ParseCard(card1) + ParseCard(card2)
	if value == 21 {
		if ParseCard(dealerCard) != 11 && ParseCard(dealerCard) != 10 {
			return "W"
		}
		return "S"
	}
	if value >= 17 && value <= 20 {
		return "S"
	}
	if value >= 12 && value <= 16 {
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	}
	return "H"
}
