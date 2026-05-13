#ifndef QORE_PQC_H
#define QORE_PQC_H

#include <stdint.h>
#include <stddef.h>

/* Dilithium-5 (ML-DSA-87) key generation.
 *
 * Buffer-length pointers serve as both IN (caller-supplied capacity) and
 * OUT (actual bytes written). Caller must pass non-null pointers to
 * pubkey_len and privkey_len with the buffer capacities; on success the
 * lib writes the actual sizes back. Returns 0 on success, negative on
 * error (see PQCError enum in libqorepqc).
 */
int32_t qore_dilithium_keygen(uint8_t *pubkey_out, size_t *pubkey_len,
                              uint8_t *privkey_out, size_t *privkey_len);

/* Dilithium-5 sign. sig_len is IN+OUT (capacity → actual). */
int32_t qore_dilithium_sign(const uint8_t *privkey, size_t privkey_len,
                            const uint8_t *message, size_t message_len,
                            uint8_t *sig_out, size_t *sig_len);

/* Dilithium-5 verify. Returns 1 on valid, 0 on invalid, negative on error. */
int32_t qore_dilithium_verify(const uint8_t *pubkey, size_t pubkey_len,
                              const uint8_t *message, size_t message_len,
                              const uint8_t *signature, size_t sig_len);

#endif
