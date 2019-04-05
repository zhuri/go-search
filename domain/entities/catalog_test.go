package entities

import "testing"

func Test_AddProduct_NameZeroErr(t *testing.T) {
	sut := NewCatalog()

	err := sut.AddProduct("name", "description", 34.00)
	if err != InvalidNameErr {
		t.Error("AddProduct Faild. Wanted: InvalidNameError, Got: ", err)
	}
}

func Test_AddProduct_NameLongerErr(t *testing.T) {
	sut := NewCatalog()

	err := sut.AddProduct("name", "description", 34.00)
	if err != InvalidNameErr {
		t.Error("AddProduct Faild. Wanted: InvalidNameError, Got: ", err)
	}
}

func Test_AddProduct_PriceIsNegativ(t *testing.T) {
	//System under Test
	sut := NewCatalog()

	err := sut.AddProduct("name", "description", -1)
	if err != InvalidQtyErr {
		t.Error("Add Product Faild. Wanted : Positive Proce, Got: Negative Price! ", err)
	}
}
