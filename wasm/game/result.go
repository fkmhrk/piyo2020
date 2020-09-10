package game

import (
	"sort"
	"time"
)

// Result describes results
type Result struct {
	StartCount int
	DeathCount int
	Days       int
	LastDate   int
	TopScores  []resultScore
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

func (r *Result) MarkPlay() {
	today := toYYYYMMDD(time.Now())
	yesterday := toYYYYMMDD(time.Now().AddDate(0, 0, -1))
	if r.LastDate == today {
		return
	}
	if r.LastDate == yesterday {
		r.LastDate = today
		r.Days++
	} else {
		r.LastDate = today
		r.Days = 1
	}
}

func (r *Result) ToMap() map[string]interface{} {
	scores := make([]map[string]interface{}, 0)
	for _, s := range r.TopScores {
		scores = append(scores, s.toMap())
	}
	return map[string]interface{}{
		"start_count": r.StartCount,
		"death_count": r.DeathCount,
		"days":        r.Days,
		"last_date":   r.LastDate,
		"scores":      scores,
	}
}

func (s *resultScore) toMap() map[string]interface{} {
	return map[string]interface{}{
		"score": s.Score,
		"date":  s.Date,
	}
}

func (r *Result) LoadFromMap(data map[string]interface{}) {
	startCount, _ := data["start_count"].(float64)
	deathCount, _ := data["death_count"].(float64)
	days, _ := data["days"].(float64)
	lastDate, _ := data["last_date"].(float64)
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
	r.StartCount = int(startCount)
	r.DeathCount = int(deathCount)
	r.Days = int(days)
	r.LastDate = int(lastDate)
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

func toYYYYMMDD(t time.Time) int {
	return t.Year()*100000 + int(t.Month())*100 + t.Day()
}
