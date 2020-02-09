package secretshare

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"secretshare/interpolation"
)

const (
	secretMax = 10000
	shareMax  = 10000
)

type shares struct {
	k  uint
	n  uint
	dp *interpolation.DataPoints
}

func contains(s []*big.Int, e *big.Int) bool {
	for _, a := range s {
		if a == nil {
			return false
		}
		if a.Cmp(e) == 0 {
			return true
		}
	}
	return false
}

func genRandom(max int64) (*big.Int, error) {
	k, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return big.NewInt(0), err
	} else {
		return big.NewInt(k.Int64()), nil
		//return float64(k.Int64()), nil
	}
}

func NewSecretshares(_k uint) (*shares, error) {
	secret, err := genRandom(secretMax)
	if err != nil {
		return nil, err
	}

	xterm := make([]*big.Int, _k)
	yterm := make([]*big.Int, _k)
	xterm[0] = big.NewInt(0)
	yterm[0] = secret

	for i := uint(1); i < _k; i++ {
		for {
			x, err := genRandom(shareMax)
			if err != nil {
				return nil, err
			}
			if !contains(xterm, x) {
				xterm[i] = x
				break
			}
		}
		for {
			y, err := genRandom(shareMax)
			if err != nil {
				return nil, err
			}
			if !contains(yterm, y) {
				yterm[i] = y
				break
			}
		}
	}
	dp, err := interpolation.NewDataPoints(xterm, yterm)
	if err != nil {
		return nil, err
	}
	n := 2*_k - 1
	sharex := make([]*big.Int, n)
	sharey := make([]*big.Int, n)
	copy(sharex, xterm[1:])
	copy(sharey, yterm[1:])

	for i := _k - 1; i < n; i++ {
		for {
			x, err := genRandom(shareMax)
			if err != nil {
				return nil, err
			}
			if !contains(sharex, x) {
				y, err := dp.CalcInterpolation(x)
				if err != nil {
					return nil, err
				}
				sharex[i] = x
				sharey[i] = y
				break
			}
		}
	}
	fmt.Println("origin secret:", secret)
	newdp, err := interpolation.NewDataPoints(sharex, sharey)
	if err != nil {
		return nil, err
	}
	return &shares{
		k:  _k,
		n:  n,
		dp: newdp,
	}, nil
}

func GetSecretKey(xterm, yterm []*big.Int) error {
	dp, err := interpolation.NewDataPoints(xterm, yterm)
	if err != nil {
		return err
	}
	secret, err := dp.CalcInterpolation(big.NewInt(0))
	if err != nil {
		return err
	}
	fmt.Println("recovered secret:", secret)
	return nil

}
