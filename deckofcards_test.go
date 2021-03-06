package deckofcards_test

import (
	"testing"

	"deckofcards"
)

func TestNew(t *testing.T) {
	deck := deckofcards.New(true)

	if !deck.Shuffled {
		t.Error("shuffled must be true")
	}

	if deck.Remaining != deckofcards.DefaultRemaining {
		t.Errorf("remaining must be %v", deckofcards.DefaultRemaining)
	}

	if deck.Cards == nil {
		t.Error("cards must not be nil")
	}
}

func TestDeck_Draw(t *testing.T) {
	var (
		count = 131
	)

	deck := deckofcards.New(false)

	deck.Draw(count)

	if deck.Remaining != 0 {
		t.Error("drawing invalid number of cards")
	}
}

func TestDeck_Delete(t *testing.T) {
	deck := deckofcards.New(false)

	deck.Delete(50)

	if deck.Remaining != deckofcards.DefaultRemaining-1 {
		t.Error("remaining must be decreasing")
	}
}

func TestDeck_Partial(t *testing.T) {
	const (
		code  = "AS"
		value = "ACE"
		suit  = "SPADES"
	)

	deck := deckofcards.New(false).Partial(code)

	if deck.Remaining != 1 {
		t.Errorf("expect 1 card, buf remaining is %v", deck.Remaining)
		t.FailNow()
	}

	card := deck.Cards[0]

	if card.Code != code {
		t.Errorf("expect %s code, but card code is %s", code, card.Code)
	}

	if card.Value != value {
		t.Errorf("expect %s value, but card value is %s", value, card.Value)

	}

	if card.Suit != suit {
		t.Errorf("expect %s suit, but card suit is %s", suit, card.Suit)
	}
}

func TestDeck_Shuffle(t *testing.T) {
	var (
		card *deckofcards.Card
	)

	deck := deckofcards.New(false)

	card = deck.Cards[0]

	deck.Shuffle()

	if deck.Cards[0].Code == card.Code {
		t.Error("deck is not shuffled")
	}
}

func TestDeck_Default(t *testing.T) {
	deck := deckofcards.New(false)

	if deck.Shuffled {
		t.Error("shuffled must be false")
	}

	if deck.Remaining != deckofcards.DefaultRemaining {
		t.Errorf("remaining must be %v", deckofcards.DefaultRemaining)
	}

	if deck.Cards == nil {
		t.Error("cards must not be nil")
	}
}

func TestDeck_Pile(t *testing.T) {
	const (
		count = 2
		name  = "player"
	)

	deck := deckofcards.New(false)

	pile := deck.Pile("player", deck.Draw(count))

	if pile.Remaining != count {
		t.Errorf("pile remaining must be %v", count)
	}

	if deck.Piles[name].Remaining != count {
		t.Errorf("pile remaining must be %v", count)
	}
}

func TestPile_Delete(t *testing.T) {
	deck := deckofcards.New(false)
	pile := deck.Pile("player", deck.Draw(1))

	pile.Delete(0)

	if pile.Remaining != 0 {
		t.Error("remaining must be 0")
	}
}

func TestPile_Draw(t *testing.T) {
	deck := deckofcards.New(false)
	pile := deck.Pile("player", deck.Draw(1))

	cards := pile.Draw(1)

	if len(cards) != 1 {
		t.Error("cards len must be 1")
	}

	if pile.Remaining != 0 {
		t.Error("remaining must be 0")
	}
}
