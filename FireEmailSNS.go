package notifier

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kkesley/s3-msgstruct"
)

//FireEmailSNS push email sns notification
func FireEmailSNS(topic string, email *msgstruct.StandardEmailStructure) error {
	if email == nil {
		return errors.New("email cannot be nil")
	}
	email.Method = "email"
	conBytes, err := json.Marshal(email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	message := string(conBytes)
	return FireSNS(topic, message)
}
