package secretshare

import (
	"math/big"
	"testing"

	//"github.com/hyunsooda/ECDSA/secp256k1/crypto"

	"gopkg.in/go-playground/assert.v1"
)

func TestOne(t *testing.T) {
	shares, err := NewSecretshares(uint(3))
	assert.Equal(t, err, nil)

	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[:3], shares.dp.Yterm[:3])
	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[1:4], shares.dp.Yterm[1:4])
	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[2:5], shares.dp.Yterm[2:5])
	// The follow test Must be same with the origin secret key
	GetSecretKey([]*big.Int{shares.dp.Xterm[0], shares.dp.Xterm[2], shares.dp.Xterm[4]},
		[]*big.Int{shares.dp.Yterm[0], shares.dp.Yterm[2], shares.dp.Yterm[4]})
	// The follow test Must be same with the origin secret key
	GetSecretKey([]*big.Int{shares.dp.Xterm[1], shares.dp.Xterm[3], shares.dp.Xterm[0]},
		[]*big.Int{shares.dp.Yterm[1], shares.dp.Yterm[3], shares.dp.Yterm[0]})
	// The follow test Must be different with the origin secret key
	GetSecretKey(shares.dp.Xterm[2:4], shares.dp.Yterm[2:4])
}

func TestTwo(t *testing.T) {
	shares, err := NewSecretshares(uint(15))
	assert.Equal(t, err, nil)

	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[:15], shares.dp.Yterm[:15])
	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[3:3+15], shares.dp.Yterm[3:3+15])
	// The follow test Must be same with the origin secret key
	GetSecretKey([]*big.Int{shares.dp.Xterm[0], shares.dp.Xterm[2], shares.dp.Xterm[4],
		shares.dp.Xterm[21], shares.dp.Xterm[23], shares.dp.Xterm[25],
		shares.dp.Xterm[13], shares.dp.Xterm[15], shares.dp.Xterm[17],
		shares.dp.Xterm[1], shares.dp.Xterm[28], shares.dp.Xterm[7],
		shares.dp.Xterm[3], shares.dp.Xterm[14], shares.dp.Xterm[11]},
		[]*big.Int{shares.dp.Yterm[0], shares.dp.Yterm[2], shares.dp.Yterm[4],
			shares.dp.Yterm[21], shares.dp.Yterm[23], shares.dp.Yterm[25],
			shares.dp.Yterm[13], shares.dp.Yterm[15], shares.dp.Yterm[17],
			shares.dp.Yterm[1], shares.dp.Yterm[28], shares.dp.Yterm[7],
			shares.dp.Yterm[3], shares.dp.Yterm[14], shares.dp.Yterm[11]})

	// The follow test Must be different with the origin secret key
	GetSecretKey(shares.dp.Xterm[5:5+9], shares.dp.Yterm[5:5+9])

}

func TestThree(t *testing.T) {
	shares, err := NewSecretshares(uint(50))
	assert.Equal(t, err, nil)

	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[:50], shares.dp.Yterm[:50])
	// The follow test Must be different with the origin secret key
	GetSecretKey(shares.dp.Xterm[:49], shares.dp.Yterm[:49])
}

func TestFour(t *testing.T) {
	shares, err := NewSecretshares(uint(100))
	assert.Equal(t, err, nil)

	// The follow test Must be same with the origin secret key
	GetSecretKey(shares.dp.Xterm[:100], shares.dp.Yterm[:100])
	// The follow test Must be different with the origin secret key
	GetSecretKey(shares.dp.Xterm[:99], shares.dp.Yterm[:99])
}
