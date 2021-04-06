package imgio

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
)

// Write takes a slice of bytes, and writes them to the current director as a
// new image file. It uses a UUID as the filename.
func Write(b []byte) error {
	id := uuid.New()
	fp := fmt.Sprintf("./%s.jpg", id.String())

	log.Printf("writing image to file: %s", fp)

	return ioutil.WriteFile(fp, b, 0600)
}
