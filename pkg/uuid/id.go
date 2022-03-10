package uuid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/oklog/ulid/v2"
	"time"
)

const RandomIdSize = 20

func RandomID() string {
	b := make([]byte, RandomIdSize)
	n, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	if n != RandomIdSize {
		panic(fmt.Sprintf("dbx: random id size is not %v", RandomIdSize))
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func ULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader).String()
}
