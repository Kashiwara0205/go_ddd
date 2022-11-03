package value_object

import (
	"errors"
)

type MailAddress struct {
	value *string
}

func NewMailAddress(value *string) (*MailAddress, error) {

	if value == nil {
		return nil, errors.New("メールが入力されていません")
	}

	mailAddress := &MailAddress{value: value}

	return mailAddress, nil
}

func (m *MailAddress) Value() string {
	return *m.value
}
