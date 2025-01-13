package main

import (
	"errors"
	"os"
)

type ProductConfig struct {
	ProductId string
	Width     float64
	Height    float64
}

func getProductConfig(productId string) (ProductConfig, error) {
	switch productId {
	case os.Getenv("ALDER_PRODUCT_ID"):
		return ProductConfig{
			ProductId: os.Getenv("ALDER_PRODUCT_ID"),
			Width:     23,
			Height:    10,
		}, nil
	case os.Getenv("FLEUR_PRODUCT_ID"):
		return ProductConfig{
			ProductId: os.Getenv("FLEUR_PRODUCT_ID"),
			Width:     15,
			Height:    11,
		}, nil
	default:
		return ProductConfig{}, errors.New("invalid product ID")
	}
}
