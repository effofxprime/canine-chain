package types

import "strconv"

func (a *ActiveDeals) IsYoung(blockHeight int64, window int64) bool {
	startBlock, err := strconv.ParseInt(a.Startblock, 10, 64)
	if err != nil {
		return false
	}

	return blockHeight <= startBlock+window
}
