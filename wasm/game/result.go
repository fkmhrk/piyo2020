package game

import (
	"sort"
)

// Result describes results
type Result struct {
	TopScores []resultScore
}

type resultScore struct {
	Score int
	Date  int64
}

// NewResult creates result object
func NewResult() *Result {
	return &Result{
		TopScores: make([]resultScore, 0, 11),
	}
}

func (r *Result) AddScore(score int, date int64) {
	rs := resultScore{
		Score: score,
		Date:  date,
	}
	r.TopScores = append(r.TopScores, rs)
	sort.Slice(r.TopScores, func(i, j int) bool {
		return r.TopScores[i].Score > r.TopScores[j].Score
	})
	if len(r.TopScores) >= 11 {
		r.TopScores = r.TopScores[:10]
	}
}

func (r *Result) ToMap() map[string]interface{} {
	scores := make([]map[string]interface{}, 0)
	for _, s := range r.TopScores {
		scores = append(scores, s.toMap())
	}
	return map[string]interface{}{
		"scores": scores,
	}
}

func (s *resultScore) toMap() map[string]interface{} {
	return map[string]interface{}{
		"score": s.Score,
		"date":  s.Date,
	}
}

func (r *Result) LoadFromMap(data map[string]interface{}) {
	scoreMaps, ok := data["scores"].([]interface{})
	if !ok {
		return
	}
	scores := make([]resultScore, 0, len(scoreMaps))
	for _, s := range scoreMaps {
		sMap, ok := s.(map[string]interface{})
		if !ok {
			continue
		}
		score := resultScore{}
		score.LoadFromMap(sMap)
		scores = append(scores, score)
	}
	r.TopScores = scores
}

func (s *resultScore) LoadFromMap(data map[string]interface{}) {
	score, ok := data["score"].(float64)
	if !ok {
		return
	}
	date, ok := data["date"].(float64)
	if !ok {
		return
	}
	s.Score = int(score)
	s.Date = int64(date)
}
