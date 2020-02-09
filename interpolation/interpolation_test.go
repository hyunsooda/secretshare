package interpolation

import (
	"fmt"
	"math/big"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestInit(t *testing.T) {
	_, err := NewDataPoints([]*big.Int{big.NewInt(0), big.NewInt(13), big.NewInt(29)},
		[]*big.Int{big.NewInt(25), big.NewInt(15), big.NewInt(7)})
	assert.Equal(t, err, nil)
}

func TestInterpolation(t *testing.T) {
	dp, err := NewDataPoints([]*big.Int{big.NewInt(0), big.NewInt(13), big.NewInt(29)},
		[]*big.Int{big.NewInt(25), big.NewInt(15), big.NewInt(7)})
	assert.Equal(t, err, nil)
	y, err := dp.CalcInterpolation(big.NewInt(17))
	fmt.Println(y)
	assert.Equal(t, err, nil)
}

//func TestMakeCurve(t *testing.T) {
//	//dp, err := NewDataPoints([]float64{-1, 1, 2, 4}, []float64{1, 2, -2, 0})
//	//dp, err := NewDataPoints([]float64{1, -1, 4, 2}, []float64{2, 1, 0, -2})
//	dp, err := NewDataPoints([]float64{0, 13, 29}, []float64{25, 15, 7})
//	assert.Equal(t, err, nil)
//	assert.Equal(t, err, nil)
//	curveFunc, err := dp.MakeCurve()
//	fmt.Println(curveFunc(17))
//	fmt.Println(curveFunc(19))
//	assert.Equal(t, err, nil)
//}
