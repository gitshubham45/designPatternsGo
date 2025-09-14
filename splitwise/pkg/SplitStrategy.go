package pkg

import "fmt"

type SplitStrategy interface {
	Split(amount float64, numUsers int, inputs []float64) ([]float64, error)
}

type EqualSplit struct{}

func (e *EqualSplit) Split(amount float64, numUsers int, inputs []float64) ([]float64, error) {
	share := amount / float64(numUsers)

	for i := 0; i < numUsers; i++ {
		inputs[i] = share
	}
	return inputs, nil
}

type ExactSplit struct{}

func (e *ExactSplit) Split(amount float64, numUsers int, inputs []float64) ([]float64, error) {
	if len(inputs) != numUsers {
		return nil, fmt.Errorf("EXACT split requires %d values", numUsers)
	}

	sum := 0.0

	for _, v := range inputs {
		sum += v
	}

	if sum != amount {
		return nil, fmt.Errorf("sum =of exact values (%f) != amount (%f)", sum, amount)
	}

	return inputs, nil
}

type PercentSplit struct{}

func (p *PercentSplit) Split(amount float64, numUsers int, inputs []float64) ([]float64, error) {
	if len(inputs) != numUsers {
		return nil, fmt.Errorf("PERCENT split requires %d values", numUsers)
	}

	sum := 0.0
	for _, v := range inputs {
		sum += float64(v)
	}

	if sum != 100 {
		return nil, fmt.Errorf("sum of percentages must be 100, got %f", sum)
	}

	values := make([]float64, numUsers)
	for i, v := range inputs {
		values[i] = amount * v / 100
	}
	return values, nil
}

func GetSplitStategy(splitType string) (SplitStrategy, error) {
	switch splitType {
	case "EQUAL":
		return &EqualSplit{}, nil
	case "EXACT":
		return &ExactSplit{}, nil
	case "PERCENT":
		return &PercentSplit{}, nil
	default:
		return nil, fmt.Errorf("unknown split type: %s", splitType)
	}
}
