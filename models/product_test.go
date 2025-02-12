package models_test

import (
	"ims/models"
	"testing"
)

func TestProductPrice(t *testing.T) {

	var p models.Product = models.Product{1, "Likhil", 23, 1, "Human"}
	p.UpdatePrice(154)
	if p.Price == 154 {
		t.Log("Running as expected")
	} else {
		t.FailNow()
	}

}
