package mmr

import (
	"fmt"
)

type Rating struct {
	Rating     float64
	Deviation  float64
	Volatility float64
	TwoSigma   float64
}

func DefaultRating() *Rating {
	return &Rating{DefaultRat, DefaultDev, DefaultVol, 2 * DefaultDev}
}

func NewRating(r, rd, s float64) *Rating {
	return &Rating{r, rd, s, 2 * rd}
}

func (r *Rating) ToGlicko2() *Rating {
	return NewRating(
		(r.Rating-DefaultRat)/glicko2Scale,
		(r.Deviation)/glicko2Scale,
		r.Volatility)
}

func (r *Rating) FromGlicko2() *Rating {
	return NewRating(
		r.Rating*glicko2Scale+DefaultRat,
		r.Deviation*glicko2Scale,
		r.Volatility)
}

func (r *Rating) String() string {
	return fmt.Sprintf("{Rating[%.3f] Deviation[%.3f] Volatility[%.3f]}",
		r.Rating, r.Deviation, r.Volatility)
}

func (r *Rating) MostlyEquals(o *Rating, epsilon float64) bool {
	return floatsMostlyEqual(r.Rating, o.Rating, epsilon) &&
		floatsMostlyEqual(r.Deviation, o.Deviation, epsilon) &&
		floatsMostlyEqual(r.Volatility, o.Volatility, epsilon)
}

func (r *Rating) Copy() *Rating {
	return &Rating{
		r.Rating,
		r.Deviation,
		r.Volatility,
		2 * r.Deviation,
	}
}
