package similarity

import (
	"errors"
	"sort"
)

// Pair is a pair of slice indexes to the sets in the input to all-pairs
// algorithms.
type Pair struct {
	X          int
	Y          int
	Similarity float64
}

type postingListEntry struct {
	setIndex      int
	tokenPosition int
	setSize       int
}

// AllPairs finds all pairs of transformed sets with similarity greater than a
// threshold.  This is an implementation of the All-Pair-Binary algorithm in the
// paper "Scaling Up All Pairs Similarity Search" by Bayardo et al., with
// position and length filter enhancement.
// Currently supported similarity functions are "jaccard" and "cosine".
// This function returns a channel of Pairs which contains the indexes to
// the input set slice.
func AllPairs(sets [][]int, similarityFunctionName string,
	similarityThreshold float64) (<-chan Pair, error) {
	if len(sets) == 0 {
		return nil, errors.New("input sets mut be a non-empty slice")
	}
	if similarityThreshold < 0 || similarityThreshold > 1.0 {
		return nil, errors.New("input similarityThreshold must be in the range [0, 1]")
	}
	var simFunc function
	if f, exists := similarityFuncs[similarityFunctionName]; exists {
		simFunc = f
	} else {
		return nil, errors.New("input similarityFunctionName does not exist")
	}
	if !symmetricSimilarityFuncs[similarityFunctionName] {
		return nil, errors.New("input similarityFunctionName is not symmetric")
	}
	overlapThresholdFunc := overlapThresholdFuncs[similarityFunctionName]
	overlapIndexThresholdFunc := overlapIndexThresholdFuncs[similarityFunctionName]
	positionFilterFunc := positionFilterFuncs[similarityFunctionName]
	pairs := make(chan Pair)
	go func() {
		// Create a slice of set indexes.
		indexes := make([]int, len(sets))
		for i := range indexes {
			indexes[i] = i
		}
		// Sort set indexes by set length.
		sort.Slice(indexes, func(i, j int) bool {
			return len(sets[i]) < len(sets[j])
		})
		defer close(pairs)
		postingLists := make(map[int][]postingListEntry)
		// Main loop of the All-Pairs algorithm.
		for _, x1 := range indexes {
			s1 := sets[x1]
			t := overlapThresholdFunc(len(s1), similarityThreshold)
			prefixSize := len(s1) - t + 1
			prefix := s1[:prefixSize]
			// Find candidates using tokens in the prefix.
			candidates := make([]int, 0)
			for p1, token := range prefix {
				for _, entry := range postingLists[token] {
					if positionFilterFunc(s1, sets[entry.setIndex], p1,
						entry.tokenPosition, similarityThreshold) {
						candidates = append(candidates, entry.setIndex)
					}
				}
			}
			// Sort and iterate through candidate indexes to verify
			// pairs.
			// TODO: optimize using partial overlaps.
			sort.Ints(candidates)
			prevCandidate := -1
			for _, x2 := range candidates {
				// Skip seen candidate.
				if x2 == prevCandidate {
					continue
				}
				prevCandidate = x2
				// Compute the exact similarity of this candidate
				sim := simFunc(s1, sets[x2])
				if sim < similarityThreshold {
					continue
				}
				if x1 > x2 {
					pairs <- Pair{x1, x2, sim}
				} else {
					pairs <- Pair{x2, x1, sim}
				}
			}
			// Insert the tokens in the prefix into index.
			t = overlapIndexThresholdFunc(len(s1), similarityThreshold)
			prefixSize = len(s1) - t + 1
			prefix = s1[:prefixSize]
			for j, token := range prefix {
				if _, exists := postingLists[token]; !exists {
					postingLists[token] = make([]postingListEntry, 0)
				}
				postingLists[token] = append(postingLists[token],
					postingListEntry{x1, j, len(sets[x1])})
			}
		}
	}()
	return pairs, nil
}
