package value_object

import (
	"errors"
	"fmt"
)

type ModelNumber struct {
	productCode *string
	branch      *string
	lot         *string
}

func NewModelNumber(productCode *string, branch *string, lot *string) (*ModelNumber, error) {

	if productCode == nil {
		return nil, errors.New("productCodeが入力されていません")
	}

	if branch == nil {
		return nil, errors.New("branchが入力されていません")
	}

	if lot == nil {
		return nil, errors.New("lotが入力されていません")
	}

	modelNumber := &ModelNumber{productCode: productCode, branch: branch, lot: lot}

	return modelNumber, nil
}

func (m *ModelNumber) ToString() string {
	return fmt.Sprintf("%s - %s - %s", *m.productCode, *m.branch, *m.lot)
}
