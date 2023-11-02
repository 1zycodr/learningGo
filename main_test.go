package main

import (
	"os"
	"testing"
)

func TestCustomSaveAndLoad(t *testing.T) {
	defer func() {
		os.Remove("_customtesting")
	}()

	entity := Custom{data: "data"}
	err := entity.Save("_customtesting")
	if err != nil {
		t.Error(err.Error())
	}

	loadedEntity := Custom{}
	err = loadedEntity.Load("_customtesting")
	if err != nil {
		t.Error(err.Error())
	}

	if entity.data != loadedEntity.data {
		t.Errorf("Loaded data: \"%v\", expected: \"data\"", loadedEntity.data)
	}
}
