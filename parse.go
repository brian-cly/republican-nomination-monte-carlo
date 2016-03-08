//
// parse.go
//

package main

import (
	"log"
	"strings"
)

func parseResponses(state string, poll Poll, responses []Responses) (trump, cruz, rubio, kasich int) {
	for _, resp := range responses {
		if resp.Choice == nil {
			log.Printf("  No Choice for %v state poll by '%v'. Skipping.\n",
				state, *poll.Pollster)
			continue
		}
		if resp.Value == nil {
			log.Printf("  No Value for %v state poll by '%v'. Skipping.\n",
				state, *poll.Pollster)
			continue
		}
		if strings.EqualFold(*resp.Choice, "trump") {
			trump = *resp.Value
		}
		if strings.EqualFold(*resp.Choice, "cruz") {
			cruz = *resp.Value
		}
		if strings.EqualFold(*resp.Choice, "rubio") {
			rubio = *resp.Value
		}
		if strings.EqualFold(*resp.Choice, "kasich") {
			kasich = *resp.Value
		}
	}
	return
}

func parseSubpopulation(state string, poll Poll, sub Subpopulations) (trump, cruz, rubio, kasich, size int) {
	if sub.Observations == nil {
		log.Printf("  No N for %v state poll by '%v'. Skipping.\n",
			state, *poll.Pollster)
		return
	}

	size = *sub.Observations
	trump, cruz, rubio, kasich = parseResponses(state, poll, sub.Responses)
	return
}

func parsePollster(poll Poll) string {
	pollster := ""
	if poll.Pollster != nil {
		pollster = *poll.Pollster
	}
	return pollster
}

func parseDateAsString(poll Poll) string {
	date := ""
	if poll.Last_updated != nil {
		date = *poll.Last_updated
	}
	return date
}

func parsePoll(state string, poll Poll) (trump, cruz, rubio, kasich, size int) {
	for _, question := range poll.Questions {
		if question.Topic != nil && strings.EqualFold(*question.Topic, "2016-president-gop-primary") {
			// given multiple subpopulations, prefer likely voters
			switch len(question.Subpopulations) {
			case 1:
				trump, cruz, rubio, kasich, size = parseSubpopulation(state, poll, question.Subpopulations[0])
			default:
				foundLikelyVoters := false
				for _, sub := range question.Subpopulations {
					if sub.Name != nil && strings.EqualFold(*sub.Name, "Likely Voters - Republican") {
						trump, cruz, rubio, kasich, size = parseSubpopulation(state, poll, sub)
						foundLikelyVoters = true
					}
				}
				if !foundLikelyVoters {
					log.Printf("  No Likely voters in multi-subpopulation poll for "+
						"%v state poll by '%v'. Skipping.\n", state, *poll.Pollster)
				}
			}
		}
	}
	return
}
