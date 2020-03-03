package chess

import (
	"joueur/base"
	"strconv"
	"strings"
)

// PlayerName should return the string name of your Player in games it plays.
func PlayerName() string {
	return "Chess Go Player"
}

/**
 * Pretty formats an FEN string to a human readable string.
 *
 * For more information on FEN (Forsyth-Edwards Notation) strings see:
 * https://wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
 */
func prettyFEN(fen string, ourColor string) string {
	// split the FEN string up to help parse it
	split := strings.Split(fen, " ")
	first := split[0] // the first part is always the board locations

	sideToMove := split[1][0] // always the second part for side to move
	usOrThem := "them"
	if sideToMove == ourColor[0] {
		usOrThem = "us"
	}

	fullmove := split[5] // always the sixth part for the full move

	lines := strings.Split(first, "/")
	sb := strings.Builder{}
	sb.WriteString("Move: " + fullmove)
	sb.WriteString("\nSide to move: " + string(sideToMove) + " (" + usOrThem + ")")
	sb.WriteString("\n   +-----------------+")

	for i, line := range lines {
		sb.WriteString("\n " + strconv.FormatInt(int64(8-i), 10) + " |")
		for _, character := range line {
			if character >= rune('0') && character <= rune('9') {
				// it is a number, so that many blank lines
				for j := 0; j < int(character-'0'); j++ {
					sb.WriteString(" .")
				}
			} else {
				sb.WriteString(" " + string(character))
			}
		}
		sb.WriteString(" |")
	}
	sb.WriteString("\n   +-----------------+\n     a b c d e f g h\n")

	return sb.String()
}

// AI is your personal AI implimentation.
type AI struct {
	base.AIImpl
	// You can add new fields here
}

// Game returns the instance of the Game this AI is currently playing.
func (ai *AI) Game() Game {
	return ai.AIImpl.Game().(Game)
}

// Player returns the instance of the Player this AI is represented by in the
// game this AI is playing.
func (ai *AI) Player() Player {
	return ai.AIImpl.Player().(Player)
}

// Start is called once the game starts and your AI has a Player and Game.
// You can initialize your AI struct here.
func (ai *AI) Start() {
	// pass
}

// GameUpdated is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai *AI) GameUpdated() {
	// pass
}

// Ended is called when the game ends, you can clean up your data and dump
// files here if need be.
func (ai *AI) Ended(won bool, reason string) {
	// pass
}

// -- Chess specific AI actions -- \\

// MakeMove this is called every time it is this AI.player's turn to make a
// move.
func (ai *AI) MakeMove() string {
	println(prettyFEN(ai.Game().Fen(), ai.Player().Color()))

	// This will only work if we are black move the pawn at b2 to b3.
	// Otherwise we will lose.
	// Your job is to code SOMETHING to parse the FEN string in some way to
	// determine a valid move, in UCI or SAN format.
	return "b2b3"
}
