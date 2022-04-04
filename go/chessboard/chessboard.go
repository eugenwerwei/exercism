package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0

	_, ok := cb[rank]

	if !ok {
		return 0
	}

	for _, v := range cb[rank] {
		if v {
			count++
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0

	if file > 8 || file < 1 {
		return 0
	}

	for _, v := range cb {
		if v[file-1] {
			count++
		}

	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		count += len(v)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		for _, v := range v {
			if v {
				count++
			}
		}
	}

	return count
}
