package lib

import "github.com/oklog/ulid/v2"

func IdsGenerateUnique() string {
	// return ulid.MustNew(ulid.Now(), nil).String()
	return ulid.Make().String()
}
