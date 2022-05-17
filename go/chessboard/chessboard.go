package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (squares int) {
	_, exists := cb[rank]
	if !exists {
		return 0
	}
	for _, b := range cb[rank] {
		if b {
			squares++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (squares int) {
	if file < 1 || file > 8 {
		return 0
	}
	for _, rank := range cb {
		if rank[file-1] {
			squares++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (squares int) {
	for _, rank := range cb {
		for range rank {
			squares++
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (squares int) {
	for _, rank := range cb {
		for _, b := range rank {
			if b {
				squares++
			}
		}
	}
	return
}
