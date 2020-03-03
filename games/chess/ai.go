package chess

import "joueur/base"

// PlayerName should return the string name of your Player in games it plays.
func PlayerName() string {
	return "Chess Go Player"
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
	// Put your game logic here for makeMove
	return ""

}
