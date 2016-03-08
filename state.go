// 
// state.go
// 

package main

import (
	"log"
	"math"
	"math/rand"
)

type StateProbability struct {
	state            string
	Trump            float64
	Cruz             float64
	N                float64
	trumpPerc        float64
	σ                float64
	TrumpProbability float64
}

// Update state data with a new poll. The new N is calculated with the actual 
// number of votes for Trump and Cruz, not the N of the poll. The effect 
// is to not count undecideds and Others. Essentially, the poll is reduced 
// to a new poll between the two potential winners. In both cases, that's
// what actually happens. Because the N is reduced, the uncertainty is 
// increased as it should be.
func (s *StateProbability) update(tPerc, cPerc, pollSize int) {
	trumpVotes := float64(tPerc) * float64(pollSize) / 100.0
	cruzVotes := float64(cPerc) * float64(pollSize) / 100.0
	s.Trump += trumpVotes
	s.Cruz += cruzVotes
	s.N += trumpVotes + cruzVotes

	s.trumpPerc = s.Trump / s.N
	s.σ = math.Sqrt((s.trumpPerc - s.trumpPerc*s.trumpPerc) / s.N)
	if min_σ != 0.0 && s.σ < min_σ {
		s.σ = min_σ
	}
	s.TrumpProbability = prOverX(0.50, s.trumpPerc, s.σ)
}

func (s *StateProbability) simulateElection(r *rand.Rand) int {
	if s.N != 0 {
		if r.Float64() < s.TrumpProbability {
			return college[s.state].votes
		}
	}
	return 0
}

func (s *StateProbability) logStateProbability() {
	if s.N != 0 {
		log.Printf("  %v: Trump polling=%6.4f, N=%d, σ=%6.4f --> Pr(Trump)=%6.4f\n",
			s.state, s.trumpPerc, int(s.N), s.σ, s.TrumpProbability)
	}
}
