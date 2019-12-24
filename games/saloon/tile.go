package saloon

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Bottle is the beer Bottle currently flying over this Tile, nil
	// otherwise.
	//
	// Value can be returned as a nil pointer.
	Bottle() Bottle

	// Cowboy is the Cowboy that is on this Tile, nil otherwise.
	//
	// Value can be returned as a nil pointer.
	Cowboy() Cowboy

	// Furnishing is the furnishing that is on this Tile, nil
	// otherwise.
	//
	// Value can be returned as a nil pointer.
	Furnishing() Furnishing

	// HasHazard is if this Tile is pathable, but has a hazard that
	// damages Cowboys that path through it.
	HasHazard() bool

	// IsBalcony is if this Tile is a balcony of the Saloon that
	// YoungGuns walk around on, and can never be pathed through by
	// Cowboys.
	IsBalcony() bool

	// TileEast is the Tile to the 'East' of this one (x+1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileEast() Tile

	// TileNorth is the Tile to the 'North' of this one (x, y-1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileNorth() Tile

	// TileSouth is the Tile to the 'South' of this one (x, y+1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileSouth() Tile

	// TileWest is the Tile to the 'West' of this one (x-1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileWest() Tile

	// X is the x (horizontal) position of this Tile.
	X() int64

	// Y is the y (vertical) position of this Tile.
	Y() int64

	// YoungGun is the YoungGun on this tile, nil otherwise.
	//
	// Value can be returned as a nil pointer.
	YoungGun() YoungGun

	// -- Tiled Game Utils -- \\

	// GetNeighbors returns an array of the neighbors of this Tile.
	GetNeighbors() []Tile

	// IsPathable returns if the Tile is pathable for FindPath
	IsPathable() bool

	// HasNeighbor checks if this Tile has a specific neighboring Tile.
	HasNeighbor(Tile) bool
}

// TileDirections are all the direction strings that Tile's neighbors can be
// in.
var TileDirections = [...]string{"North", "South", "East", "West"}
