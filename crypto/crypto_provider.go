package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// Sign signs arbitrary data using ECDSA.
func Sign(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	// hash message
	digest := sha256.Sum256(data)

	// sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, privKey, digest[:])
	if err != nil {
		return nil, err
	}

	// encode the signature {R, S}
	// big.Int.Bytes() will need padding in the case of leading zero bytes
	params := privKey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// Verify checks a raw ECDSA signature.
// Returns true if it's valid and false if not.
func Verify(digest []byte, signature []byte, pubKey *ecdsa.PublicKey) bool {

	curveOrderByteSize := pubKey.Curve.Params().P.BitLen() / 8

	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])

	return ecdsa.Verify(pubKey, digest[:], r, s)
}

// Verify as sign was created by address owner
func IsSigned(address string, msg string, r string, s string, v int) (bool, error) {
	// prepare
	addressAsBytes, _ := hex.DecodeString(address)
	msgAsBytes, _ := hex.DecodeString(msg)
	rAsBytes, _ := hex.DecodeString(r)
	sAsBytes, _ := hex.DecodeString(s)
	sign := make([]byte, 64)
	copy(sign[32-len(rAsBytes):32], rAsBytes)
	copy(sign[64-len(sAsBytes):64], sAsBytes)

	key1, key2, err := ecRecovery(msgAsBytes, sign)
	if err != nil {
		return false, err
	}
	result := Verify(msgAsBytes, sign, key1)
	if result == false {
		return false, nil
	}
	restoreAddress := AddressFromPubKey(key1)
	if bytes.Compare(addressAsBytes, restoreAddress) == 0 {
		return true, nil
	}
	result = Verify(msgAsBytes, sign, key2)
	if result == false {
		return false, nil
	}
	restoreAddress = AddressFromPubKey(key2)
	return bytes.Compare(addressAsBytes, restoreAddress) == 0, nil
}

// Create account address from Public Key
func AddressFromPubKey(pubKey *ecdsa.PublicKey) []byte {
	pubKeyAsBytes := compressPubKey(pubKey)
	return AddressFromPubkeyBytes(pubKeyAsBytes)
}

// Create account address from Public Key sa bytes
func AddressFromPubkeyBytes(pubKey []byte) []byte {
	hash := ripemd160Hash(pubKey)
	return hash[len(hash)-20:]
}

// CompressPubKey encodes a public key to the 64-byte compressed format.
func compressPubKey(pubKey *ecdsa.PublicKey) []byte {
	params := pubKey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8

	pubKeyAsBytes := make([]byte, curveOrderByteSize*2)
	pubKeyXBytes := pubKey.X.Bytes()
	pubKeyYBytes := pubKey.Y.Bytes()

	copy(pubKeyAsBytes[(curveOrderByteSize-len(pubKeyXBytes)):curveOrderByteSize], pubKeyXBytes)
	copy(pubKeyAsBytes[curveOrderByteSize+(curveOrderByteSize-len(pubKeyYBytes)):], pubKeyYBytes)

	return pubKeyAsBytes
}

func ripemd160Hash(data []byte) []byte {
	digest := sha256.Sum256(data)
	hasher := ripemd160.New()
	hasher.Write(digest[:])
	hash := hasher.Sum(nil)
	return hash[:]
}

func ecRecovery(data []byte, rawSign []byte) (*ecdsa.PublicKey, *ecdsa.PublicKey, error) {
	r := big.Int{}
	s := big.Int{}
	sigLen := len(rawSign)
	r.SetBytes(rawSign[:(sigLen / 2)])
	s.SetBytes(rawSign[(sigLen / 2):])

	expy := new(big.Int).Sub(elliptic.P256().Params().N, big.NewInt(2))
	rinv := new(big.Int).Exp(&r, expy, elliptic.P256().Params().N)
	z := new(big.Int).SetBytes(data)

	xxx := new(big.Int).Mul(&r, &r)
	xxx.Mul(xxx, &r)

	ax := new(big.Int).Mul(big.NewInt(3), &r)

	yy := new(big.Int).Sub(xxx, ax)
	yy.Add(yy, elliptic.P256().Params().B)

	y1 := new(big.Int).ModSqrt(yy, elliptic.P256().Params().P)
	if y1 == nil {
		return nil, nil, fmt.Errorf("can not revcovery public key")
	}

	y2 := new(big.Int).Neg(y1)
	y2.Mod(y2, elliptic.P256().Params().P)
	p1, p2 := elliptic.P256().ScalarMult(&r, y1, s.Bytes())
	p3, p4 := elliptic.P256().ScalarBaseMult(z.Bytes())

	p5 := new(big.Int).Neg(p4)
	p5.Mod(p5, elliptic.P256().Params().P)

	q1, q2 := elliptic.P256().Add(p1, p2, p3, p5)
	q3, q4 := elliptic.P256().ScalarMult(q1, q2, rinv.Bytes())

	n1, n2 := elliptic.P256().ScalarMult(&r, y2, s.Bytes())
	n3, n4 := elliptic.P256().ScalarBaseMult(z.Bytes())

	n5 := new(big.Int).Neg(n4)
	n5.Mod(n5, elliptic.P256().Params().P)

	q5, q6 := elliptic.P256().Add(n1, n2, n3, n5)
	q7, q8 := elliptic.P256().ScalarMult(q5, q6, rinv.Bytes())

	key1 := ecdsa.PublicKey{Curve: elliptic.P256(), X: q3, Y: q4}
	key2 := ecdsa.PublicKey{Curve: elliptic.P256(), X: q7, Y: q8}
	return &key1, &key2, nil
}
