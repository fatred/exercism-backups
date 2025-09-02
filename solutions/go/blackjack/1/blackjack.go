package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    value := 0
    switch {
        case card == "ace": value = 11
    	case card == "two": value = 2
    	case card == "three": value = 3
    	case card == "four": value = 4
    	case card == "five": value = 5
    	case card == "six": value = 6
    	case card == "seven": value = 7
    	case card == "eight": value = 8
    	case card == "nine": value = 9
    	case card == "ten": value = 10
		case card == "jack": value = 10
    	case card == "queen": value = 10
    	case card == "king": value = 10
    }
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    card1Value := ParseCard(card1)
    card2Value := ParseCard(card2)
    dealerCardValue := ParseCard(dealerCard)

    action := "W"
    
    if card1Value == 11 && card2Value == 11 {
        action = "P"
    } else {
    	switch {
            case (card1Value + card2Value == 21) && (dealerCardValue < 10): action = "W"
        	case (card1Value + card2Value == 21) && (dealerCardValue >= 10): action = "S"
        	case (card1Value + card2Value >= 17) && (card1Value + card2Value <= 20): action = "S"
        	case ((card1Value + card2Value >= 12) && (card1Value + card2Value <= 16)) && (dealerCardValue < 7): action = "S"
        	case ((card1Value + card2Value >= 12) && (card1Value + card2Value <= 16)) && (dealerCardValue >= 7): action = "H"
	        case (card1Value + card2Value <= 11): action = "H"
        }
    }
	return action
}
