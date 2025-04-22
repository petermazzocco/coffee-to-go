package bufferHandler

import (
	"bytes"
	"fmt"
	"os"
)

func SecretItem(code string) error {
	var b bytes.Buffer
	if code == "secret" {
		b.Write([]byte(code + " "))
	} else {
		return fmt.Errorf("Secret denied! You cannot access the secret item on the secret menu")
	}
	fmt.Fprintf(&b, "item is unlocked!")
	b.WriteTo(os.Stdout)
	return nil
}
