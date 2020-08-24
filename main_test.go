package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCsvBayesianGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CsvBayesianGo Suite")
}
