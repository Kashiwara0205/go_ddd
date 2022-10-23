package ch02

import (
	"testing"
)

func TestNewModelNumber(t *testing.T) {
	inputProductCode := "productCode"
	inputBranch := "branch"
	inputLot := "lot"

	modelNumber, _ := NewModelNumber(&inputProductCode, &inputBranch, &inputLot)

	if *modelNumber.productCode != "productCode" {
		t.Errorf("productCodeがvalueに入っていない")
	}

	if *modelNumber.branch != "branch" {
		t.Errorf("branchがvalueに入っていない")
	}

	if *modelNumber.lot != "lot" {
		t.Errorf("lotがvalueに入っていない")
	}
}

func TestProductCodeIsNil(t *testing.T) {
	inputBranch := "branch"
	inputLot := "lot"

	_, err := NewModelNumber(nil, &inputBranch, &inputLot)

	if err.Error() != "productCodeが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestBranchIsNil(t *testing.T) {
	inputProductCode := "productCode"
	inputLot := "lot"

	_, err := NewModelNumber(&inputProductCode, nil, &inputLot)

	if err.Error() != "branchが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestBranchIsLot(t *testing.T) {
	inputProductCode := "productCode"
	inputBranch := "branch"

	_, err := NewModelNumber(&inputProductCode, &inputBranch, nil)

	if err.Error() != "lotが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestModelNumberToString(t *testing.T) {
	inputProductCode := "productCode"
	inputBranch := "branch"
	inputLot := "lot"

	modelNumber, _ := NewModelNumber(&inputProductCode, &inputBranch, &inputLot)

	if modelNumber.ToString() != "productCode - branch - lot" {
		t.Errorf("意図しない値が出力されている %s", modelNumber.ToString())
	}
}
