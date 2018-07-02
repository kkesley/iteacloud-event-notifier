package notifier

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kkesley/s3-msgstruct"
)

//FireSMSSNS push sms sns notification
func FireSMSSNS(topic string, sms *msgstruct.StandardSMSStructure) error {
	if sms == nil {
		return errors.New("sms cannot be nil")
	}
	sms.Method = "sms"
	conBytes, err := json.Marshal(sms)
	if err != nil {
		fmt.Println(err)
		return err
	}
	message := string(conBytes)
	return FireSNS(topic, message)
}
