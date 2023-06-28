package passwordHash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrirInvalidHash       = errors.New("The encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("Incompatible varsion of argon2")
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func GenerateFromPassword(password string) (encodedHash string, err error) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	// generate a cryptograpic sure random salt
	salt, err := generateRandomBytes(p.saltLength)
	fmt.Println("generated random bytes", salt)
	if err != nil {
		return "", err
	}
	// pass the plaintext password, salt and params to argon2.Idkey
	// generate a hash of the password using argon2id
	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash = fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.memory,
		p.iterations,
		p.parallelism,
		b64Salt,
		b64Hash,
	)
	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := DecodedHash(encodedHash)
	if err != nil {
		return false, err
	}
	otherHash := argon2.IDKey(
		[]byte(password),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength,
	)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func DecodedHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrirInvalidHash
	}
	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}
	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}
	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))
	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

// generate a random salt from crypto/rand package to use for agron2
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// argon2 algorithm concept
// memory - the ammount of memory use by the algorithm
// iterations -the number of iteration of passover memory
// paralleism -the number of threads use by the algorithm
// salt length -length of the random salt 16 bytes is recommanded
// key length - the length of the generated key (pass word hash) 16b or more
