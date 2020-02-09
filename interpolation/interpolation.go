package interpolation

import (
	"errors"
	"math/big"
)

var (
	// P is prime of standard secp256k1 parameter
	P, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
)

type DataPoints struct {
	Xterm []*big.Int
	Yterm []*big.Int
}

func NewDataPoints(xterm, yterm []*big.Int) (*DataPoints, error) {
	if len(xterm) != len(yterm) {
		return nil, errors.New("Not same length between xterm and yterm")
	}
	return &DataPoints{
		Xterm: xterm,
		Yterm: yterm,
	}, nil
}

func (d *DataPoints) checkParams() error {
	checkParams := len(d.Xterm) == len(d.Yterm)
	if checkParams == false {
		return errors.New("parameters are invalid")
	}
	return nil
}

// calcInterpolation calculate f(x) given x by Lagrange interpolation
func (d *DataPoints) CalcInterpolation(newx *big.Int) (*big.Int, error) {
	if err := d.checkParams(); err != nil {
		return big.NewInt(0), err
	}
	var product *big.Int
	var sum = big.NewInt(0)
	n := len(d.Xterm)

	for i := 0; i < n; i++ {
		product = d.Yterm[i]
		for j := 0; j < n; j++ {
			if i != j {
				//product = product * (newx - d.Xterm[j]) / (d.Xterm[i] - d.Xterm[j])
				xm := new(big.Int).Sub(newx, d.Xterm[j])
				xm = new(big.Int).Mod(xm, P)

				product = new(big.Int).Mul(product, xm)
				product = new(big.Int).Mod(product, P)

				dxm := new(big.Int).Sub(d.Xterm[i], d.Xterm[j])
				dxm = new(big.Int).Mod(dxm, P)

				pminus2 := new(big.Int).Sub(P, big.NewInt(2))
				dxm = new(big.Int).Exp(dxm, pminus2, P) // (d.Xterm[i] - d.Xterm[j])^-1

				product = new(big.Int).Mul(product, dxm)
				product = new(big.Int).Mod(product, P)
			}
		}
		sum = new(big.Int).Add(sum, product)
		sum = new(big.Int).Mod(sum, P)
		//sum += product
	}
	return sum, nil
}

/*
// for float64 version, not the type of big.int
func (d *DataPoints) MakeCurve() (func(float64) float64, error) {
	if err := d.checkParams(); err != nil {
		return nil, err
	}

	n := len(d.Yterm)
	x := make([]float64, n)
	y := make([][]float64, n)
	copy(x, d.Xterm)

	for i := 0; i < n; i++ {
		y[i] = make([]float64, n)
		y[i][0] = d.Yterm[i]
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			y[i][j] = (y[i+1][j-1] - y[i][j-1]) / (x[i+j] - x[i])
		}
	}
	return func(newx float64) float64 {
		var nthTerm float64
		yint := make([]float64, n)

		nthTerm = 1
		yint[0] = y[0][0]
		for i := 1; i < n; i++ {
			nthTerm = nthTerm * (newx - x[i-1])
			yint[i] = yint[i-1] + y[0][i]*nthTerm

		}
		return yint[len(yint)-1]
	}, nil
}
*/
