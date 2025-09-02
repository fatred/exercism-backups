package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
    for _, place := range cb[file] {
        if place {
            count++
        }
    }
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 0 && rank > 8 {
		return 0
	} else {
		for oi := range cb {
			for ii, place := range cb[oi] {
				if ii == rank-1 && place {
					count++
				}
			}
		}
		return count
	}
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	h := len(cb)
    v := len(cb["A"])
    return h*v
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
    for oi := range cb {
        for _, place := range cb[oi] {
            if place {
                count++
            }
        }
    }
	return count
}
