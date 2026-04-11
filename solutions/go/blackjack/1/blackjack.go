package blackjack

func ParseCard(card string) int {
	value := 0
    switch {
        case card == "ace":
        	value = 11
        case card == "two":
        	value = 2
        case card == "three":
        	value = 3
        case card == "four":
        	value = 4
        case card == "five":
        	value = 5
        case card == "six":
        	value = 6
        case card == "seven":
        	value = 7
        case card == "eight":
        	value = 8
        case card == "nine":
        	value = 9
        case card == "ten" || card == "jack" || card == "queen" || card == "king":
        	value = 10
    }
    return value
}


func FirstTurn(card1, card2, dealerCard string) string {
    var decision string
    cardsSum := ParseCard(card1) + ParseCard(card2)
    parsedDealerCard := ParseCard(dealerCard)
    switch {
        case card1 == "ace" && card2 == "ace":
        	decision = "P"
    	case cardsSum == 21 && parsedDealerCard < 10:
        	decision = "W"
        case cardsSum >= 17:
        	decision = "S"
        case cardsSum >= 12 && parsedDealerCard >= 7:
        	decision = "H"
        case cardsSum >= 12:
        	decision = "S"
        case cardsSum >= 12:
        	decision = "S"
        default:
        	decision = "H"
    }
    return decision
}
