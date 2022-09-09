package mmr

import (
	"errors"
	"log"
	"math"
)

var (
	DefaultTau = 0.3
	DefaultRat = 1500.0
	DefaultDev = 350.0
	DefaultVol = 0.06
)

const (
	piSq         = math.Pi * math.Pi
	glicko2Scale = 173.7178
)

type Result float64

const (
	Win  Result = 1
	Loss Result = 0
	Draw Result = 0.5
)

// Important: players array must be sorted in order of game rank
// Assumes there cannot be ties (if people have same accuracy, whoever came earlier must win)
func CalculateMatchMMRs(ratings []*Rating) ([]*Rating, error) {
  var err error
  for i, p := range(ratings) {
    res := genRatings(len(ratings), i)
    o := append(ratings[:i], ratings[i+1:]...)
    ratings[i], err = CalculateRating(p, o, res)
    if err != nil {
      return nil, err
    }
  }
  return ratings, nil
}

func genRatings(num int, pos int) []Result {
  res := make([]Result, num - 1)
  for i := range(res) {
    if i >= pos {
      res[i] = Win
    } else {
      res[i] = Loss
    }
  }
  return res
}

func newVolatility(estVar, estImp float64, p *Rating) float64 {
	epsilon := 0.000001
	a := math.Log(sq(p.Volatility))
	deltaSq := sq(estImp)
	phiSq := sq(p.Deviation)
	tauSq := sq(DefaultTau)
	maxIter := 100
	f := func(x float64) float64 {
		eX := math.Exp(x)
		return eX*(deltaSq-phiSq-estVar-eX)/
			(2*sq(phiSq+estVar+eX)) - (x-a)/tauSq
	}
	A := a
	B := 0.0
	if deltaSq > (phiSq + estVar) {
		B = math.Log(deltaSq - phiSq - estVar)
	} else {
		val := -1.0
		k := 1
		for ; val < 0; k++ {
			val = f(a - float64(k)*DefaultTau)
		}
		B = a - float64(k)*DefaultTau
	}
	fA := f(A)
	fB := f(B)
	fC := 0.0
	iter := 0
	for math.Abs(B-A) > epsilon && iter < maxIter {
		C := A + (A-B)*fA/(fB-fA)
		fC = f(C)
		if fC*fB < 0 {
			A = B
			fA = fB
		} else {
			fA = fA / 2
		}
		B = C
		fB = fC
		iter++
	}
	if iter == maxIter-1 {
		log.Println("Fall through! Too many iterations")
	}
	newVol := math.Exp(A / 2)
	return newVol
}

func floatsMostlyEqual(v1, v2, epsilon float64) bool {
	return math.Abs(v1-v2) < epsilon
}

func sq(x float64) float64 {
	return x * x
}

func ee(r, ri, devi float64) float64 {
	return 1.0 / (1 + math.Exp(-gee(devi)*(r-ri)))
}

func gee(dev float64) float64 {
	return 1 / math.Sqrt(1+3*dev*dev/piSq)
}

func newDeviation(dev, newVol, estVar float64) float64 {
	phip := math.Sqrt(dev*dev + newVol*newVol)
	return 1.0 / math.Sqrt(1.0/(phip*phip)+1.0/(estVar))
}

func newRatingVal(oldRating, newDev, estImpPart float64) float64 {
	return oldRating + newDev*newDev*estImpPart
}

func CalculateRating(player *Rating, opponents []*Rating, res []Result) (*Rating, error) {
	if len(opponents) != len(res) {
		return nil, errors.New("number of opponents must == number of results")
	}

	p2 := player.ToGlicko2()
	gees := make([]float64, len(opponents))
	ees := make([]float64, len(opponents))
	for i := range opponents {
		o := opponents[i].ToGlicko2()
		gees[i] = gee(o.Deviation)
		ees[i] = ee(p2.Rating, o.Rating, o.Deviation)
	}

	estVar := estVariance(gees, ees)
	estImpPart := estImprovePartial(gees, ees, res)
	estImp := estVar * estImpPart
	newVol := newVolatility(estVar, estImp, p2)
	newDev := newDeviation(p2.Deviation, newVol, estVar)
	newRating := newRatingVal(p2.Rating, newDev, estImpPart)
	rt := NewRating(newRating, newDev, newVol).FromGlicko2()

	if rt.Deviation > DefaultDev {
		rt.Deviation = DefaultDev
	}

	return rt, nil
}

func estVariance(gees, ees []float64) float64 {
	out := 0.0
	for i := range gees {
		out += sq(gees[i]) * ees[i] * (1 - ees[i])
	}
	return 1.0 / out
}

func estImprovePartial(gees, ees []float64, r []Result) float64 {
	out := 0.0
	for i := range gees {
		out += gees[i] * (float64(r[i]) - ees[i])
	}
	return out
}
