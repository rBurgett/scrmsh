package service

import (
	"github.com/google/uuid"
	"github.com/rBurgett/scrmsh/internal/constants"
	"strings"
)

type Game struct {
	ID            string
	CurrentPlayer string
	Moves         []Move
	Owner         string
	Players       []Player
	Status        constants.GameStatus
}

func (g *Game) ExecuteMove(playerID string, playerCardPosition int, targetID string, targetCardPosition int) (Move, error) {
	if playerID != g.CurrentPlayer {
		return Move{}, constants.ErrorWrongUserTurn
	}

	p, err := g.GetPlayer(playerID)
	if err != nil {
		return Move{}, err
	}
	pc, err := p.GetCard(playerCardPosition)
	if err != nil {
		return Move{}, err
	}

	t, err := g.GetPlayer(targetID)
	if err != nil {
		return Move{}, err
	}
	tc, err := t.GetCard(targetCardPosition)
	if err != nil {
		return Move{}, err
	}

	m := Move{
		Player:                   playerID,
		PlayerCardPosition:       playerCardPosition,
		PlayerCardType:           pc,
		TargetPlayer:             targetID,
		TargetPlayerCardPosition: targetCardPosition,
		TargetPlayerCardType:     tc,
	}

	winner, err := DetermineMoveWinner(m)
	if err != nil {
		return Move{}, err
	}

	m.Winner = winner

	// remove card(s) from losers
	// handle a win

	g.Moves = append(g.Moves, m)

	return m, nil
}

func (g *Game) GetPlayer(id string) (Player, error) {
	for _, p := range g.Players {
		if p.ID == id {
			return p, nil
		}
	}

	return Player{}, constants.ErrorUserNotFound
}

func CreateGame(owner Player) (Game, error) {
	owner.ID = strings.TrimSpace(owner.ID)
	owner.Name = strings.TrimSpace(owner.Name)
	owner.Secret = strings.TrimSpace(owner.Secret)
	if err := owner.Validate(); err != nil {
		return Game{}, err
	}
	owner.Status = constants.PlayerStatusAccepted

	return Game{
		ID:            uuid.NewString(),
		CurrentPlayer: "",
		Owner:         owner.ID,
		Players:       []Player{owner},
		Status:        constants.GameStatusOpen,
	}, nil
}

type Move struct {
	Player                   string
	PlayerCardPosition       int
	PlayerCardType           constants.CardType
	TargetPlayer             string
	TargetPlayerCardPosition int
	TargetPlayerCardType     constants.CardType
	Winner                   string
}

func DetermineMoveWinner(m Move) (string, error) {
	switch m.PlayerCardType {
	case constants.CardTypeDagger:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return "", nil
		case constants.CardTypeShortSword:
			return m.TargetPlayer, nil
		case constants.CardTypeMace:
			return m.TargetPlayer, nil
		case constants.CardTypeBattleAxe:
			return m.TargetPlayer, nil
		case constants.CardTypeSpear:
			return m.TargetPlayer, nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeShortSword:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return "", nil
		case constants.CardTypeMace:
			return m.TargetPlayer, nil
		case constants.CardTypeBattleAxe:
			return m.TargetPlayer, nil
		case constants.CardTypeSpear:
			return m.TargetPlayer, nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeMace:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return m.Player, nil
		case constants.CardTypeMace:
			return "", nil
		case constants.CardTypeBattleAxe:
			return m.TargetPlayer, nil
		case constants.CardTypeSpear:
			return m.TargetPlayer, nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeBattleAxe:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return m.Player, nil
		case constants.CardTypeMace:
			return m.Player, nil
		case constants.CardTypeBattleAxe:
			return "", nil
		case constants.CardTypeSpear:
			return m.TargetPlayer, nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeSpear:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return m.Player, nil
		case constants.CardTypeMace:
			return m.Player, nil
		case constants.CardTypeBattleAxe:
			return m.Player, nil
		case constants.CardTypeSpear:
			return "", nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeLongSword:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return m.Player, nil
		case constants.CardTypeMace:
			return m.Player, nil
		case constants.CardTypeBattleAxe:
			return m.Player, nil
		case constants.CardTypeSpear:
			return m.Player, nil
		case constants.CardTypeLongSword:
			return "", nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeArcher:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.Player, nil
		case constants.CardTypeShortSword:
			return m.Player, nil
		case constants.CardTypeMace:
			return m.Player, nil
		case constants.CardTypeBattleAxe:
			return m.Player, nil
		case constants.CardTypeSpear:
			return m.Player, nil
		case constants.CardTypeLongSword:
			return m.Player, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return m.Player, nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	case constants.CardTypeShield:
		return m.Player, nil
	case constants.CardTypeCrown:
		switch m.TargetPlayerCardType {
		case constants.CardTypeDagger:
			return m.TargetPlayer, nil
		case constants.CardTypeShortSword:
			return m.TargetPlayer, nil
		case constants.CardTypeMace:
			return m.TargetPlayer, nil
		case constants.CardTypeBattleAxe:
			return m.TargetPlayer, nil
		case constants.CardTypeSpear:
			return m.TargetPlayer, nil
		case constants.CardTypeLongSword:
			return m.TargetPlayer, nil
		case constants.CardTypeArcher:
			return m.Player, nil
		case constants.CardTypeShield:
			return "", nil
		case constants.CardTypeCrown:
			return m.Player, nil
		}
	}

	return "", constants.ErrorIllegalMove
}
