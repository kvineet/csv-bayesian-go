package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kvineet/csv-bayesian-go"
)

var _ = Describe("Utils", func() {
	Describe("Unique", func() {
		var input, expected []string

		Describe("signature", func() {
			BeforeEach(func() {
				input = []string{"test"}
			})
			It("should return slice of strings", func() {
				output := Unique(input...)
				Expect(output).Should(BeAssignableToTypeOf(expected))
			})
		})

		Describe("Behaviour", func() {

			When("unique elements are passed", func() {

				BeforeEach(func() {
					input = []string{"elem1", "elem2", "elem3"}
					expected = []string{"elem1", "elem2", "elem3"}
				})
				It("should return same elements in slice", func() {
					output := Unique(input...)
					Expect(output).Should(Equal(expected))
				})
			})
			When("elements with duplicates are passed", func() {
				BeforeEach(func() {
					input = []string{"elem1", "elem2", "elem1"}
					expected = []string{"elem1", "elem2"}

				})
				It("should return unique elements without duplicates", func() {
					output := Unique(input...)
					Expect(output).Should(Equal(expected))
				})
			})
		})
	})
	Describe("FloatEquals", func() {
		var inputa, inputb float64
		var expected bool

		Describe("signature", func() {
			BeforeEach(func() {
				inputa = float64(1) / float64(3)
				inputb = float64(1) / float64(5)
			})
			It("should return slice of strings", func() {
				output := FloatEquals(inputa, inputb)
				Expect(output).Should(BeAssignableToTypeOf(expected))
			})
		})
		Describe("behaviour", func() {
			When("two unequal floats are passed", func() {
				It("should return false", func() {
					inputa = float64(1) / float64(3)
					inputb = float64(1) / float64(5)
					expected = false
					output := FloatEquals(inputa, inputb)
					Expect(output).To(Equal(expected))
				})
			})
			When("two same floats are passed", func() {
				It("should return true", func() {
					inputa = float64(1) / float64(3)
					expected = true
					output := FloatEquals(inputa, inputa)
					Expect(output).To(Equal(expected))
				})
			})
			When("two floats with same value are passed", func() {
				It("should return true", func() {
					inputa = float64(1) / float64(3)
					inputb = float64(1) / float64(3)
					expected = true
					output := FloatEquals(inputa, inputb)
					Expect(output).To(Equal(expected))
				})

			})
			When("two floats with same value are passed", func() {
				It("should return true", func() {
					inputa = float64(0.3333333333333333333333333)
					inputb = float64(0.33333333333333)
					expected = true
					output := FloatEquals(inputa, inputb)
					Expect(output).To(Equal(expected))
				})

			})
		})
	})
})
