//go:build cgo

package pqc

/*
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/../../lib/darwin_arm64 -lqorepqc
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/../../lib/darwin_amd64 -lqorepqc
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/../../lib/linux_amd64 -lqorepqc
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/../../lib/linux_arm64 -lqorepqc

#include "pqc.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Dilithium-5 (ML-DSA-87) constant sizes. These match the libqorepqc
// canonical sizes; the lib's keygen returns these via the OUT length
// pointers, but the constants give us tight upfront buffer sizing.
const (
	DilithiumPublicKeySize  = 2592
	DilithiumPrivateKeySize = 4896
	DilithiumSignatureSize  = 4627
)

// DilithiumKeygen generates a new Dilithium-5 keypair.
//
// The libqorepqc FFI is:
//
//	int32_t qore_dilithium_keygen(uint8_t *pubkey_out, size_t *pubkey_len,
//	                              uint8_t *privkey_out, size_t *privkey_len);
//
// pubkey_len and privkey_len are IN+OUT: callers pass the buffer capacity,
// the lib writes the actual key size back. Returns 0 on success.
func DilithiumKeygen() (pubkey []byte, privkey []byte, err error) {
	pk := make([]byte, DilithiumPublicKeySize)
	sk := make([]byte, DilithiumPrivateKeySize)
	pkLen := C.size_t(len(pk))
	skLen := C.size_t(len(sk))

	rc := C.qore_dilithium_keygen(
		(*C.uint8_t)(unsafe.Pointer(&pk[0])),
		&pkLen,
		(*C.uint8_t)(unsafe.Pointer(&sk[0])),
		&skLen,
	)
	if rc != 0 {
		return nil, nil, fmt.Errorf("dilithium keygen failed: code %d", int32(rc))
	}
	return pk[:pkLen], sk[:skLen], nil
}

// DilithiumSign signs a message with a Dilithium-5 private key.
//
// The libqorepqc FFI is:
//
//	int32_t qore_dilithium_sign(const uint8_t *privkey, size_t privkey_len,
//	                            const uint8_t *message, size_t message_len,
//	                            uint8_t *sig_out, size_t *sig_len);
//
// sig_len is IN+OUT (capacity → actual). Returns 0 on success.
func DilithiumSign(privkey, message []byte) ([]byte, error) {
	if len(privkey) != DilithiumPrivateKeySize {
		return nil, fmt.Errorf("invalid private key size: got %d, want %d",
			len(privkey), DilithiumPrivateKeySize)
	}
	if len(message) == 0 {
		return nil, fmt.Errorf("dilithium sign: message must not be empty")
	}

	sig := make([]byte, DilithiumSignatureSize)
	sigLen := C.size_t(len(sig))

	rc := C.qore_dilithium_sign(
		(*C.uint8_t)(unsafe.Pointer(&privkey[0])),
		C.size_t(len(privkey)),
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&sig[0])),
		&sigLen,
	)
	if rc != 0 {
		return nil, fmt.Errorf("dilithium sign failed: code %d", int32(rc))
	}
	return sig[:sigLen], nil
}

// DilithiumVerify verifies a Dilithium-5 signature.
//
// The libqorepqc FFI is:
//
//	int32_t qore_dilithium_verify(const uint8_t *pubkey, size_t pubkey_len,
//	                              const uint8_t *message, size_t message_len,
//	                              const uint8_t *signature, size_t sig_len);
//
// Returns 1 on valid, 0 on invalid, negative on error. (NOT a standard
// success=0 convention — this is the canonical 'verify result' shape.)
func DilithiumVerify(pubkey, message, signature []byte) (bool, error) {
	if len(pubkey) != DilithiumPublicKeySize {
		return false, fmt.Errorf("invalid public key size: got %d, want %d",
			len(pubkey), DilithiumPublicKeySize)
	}
	if len(signature) == 0 {
		return false, fmt.Errorf("dilithium verify: signature must not be empty")
	}
	if len(message) == 0 {
		return false, fmt.Errorf("dilithium verify: message must not be empty")
	}

	rc := C.qore_dilithium_verify(
		(*C.uint8_t)(unsafe.Pointer(&pubkey[0])),
		C.size_t(len(pubkey)),
		(*C.uint8_t)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uint8_t)(unsafe.Pointer(&signature[0])),
		C.size_t(len(signature)),
	)
	switch {
	case rc == 1:
		return true, nil
	case rc == 0:
		return false, nil
	default:
		return false, fmt.Errorf("dilithium verify failed: code %d", int32(rc))
	}
}
