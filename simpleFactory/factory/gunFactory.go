package factory

import (
	"fmt"
)

func GetGun(gunType string) (IGun, error) {
	if gunType == "AK47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
