//go:build !cgo

package pqc

import "fmt"

const (
	DilithiumPublicKeySize  = 2592
	DilithiumPrivateKeySize = 4896
	DilithiumSignatureSize  = 4627
)

// DilithiumKeygen returns an error in non-CGO builds. The PQC FFI bridge
// requires a C toolchain and the libqorepqc native library, both of which
// are unavailable on Windows builds and on cross-compiles where CGO is
// disabled. Operators on those platforms cannot generate or use Dilithium-5
// keys; everything else (telemetry, dashboard, light-client header sync)
// works normally.
func DilithiumKeygen() (pubkey []byte, privkey []byte, err error) {
	return nil, nil, fmt.Errorf("dilithium keygen unavailable: build was compiled without CGO")
}

// DilithiumSign returns an error in non-CGO builds.
func DilithiumSign(privkey, message []byte) ([]byte, error) {
	return nil, fmt.Errorf("dilithium sign unavailable: build was compiled without CGO")
}

// DilithiumVerify returns an error in non-CGO builds.
func DilithiumVerify(pubkey, message, signature []byte) (bool, error) {
	return false, fmt.Errorf("dilithium verify unavailable: build was compiled without CGO")
}
