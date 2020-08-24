package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kvineet/csv-bayesian-go"
)

var _ = Describe("classifer", func() {

	Describe("initialize", func() {
		Context("from classNames", func() {
			Describe("negative test cases", func() {
				It("should error for empty classes", func() {
					classifier, err := NewClassifier()
					Expect(err).Should(Equal(ErrNotEnoughClassesDefined))
					Expect(classifier).To(BeNil())
				})
				It("should error for less than 2 classes", func() {
					classifier, err := NewClassifier("class1")
					Expect(err).Should(Equal(ErrNotEnoughClassesDefined))
					Expect(classifier).To(BeNil())
				})
				It("should error for less than 2 unique classes", func() {

					classifier, err := NewClassifier("class1", "class1", "class1")
					Expect(err).Should(Equal(ErrNotEnoughClassesDefined))
					Expect(classifier).To(BeNil())
				})
				It("should error for empty class name", func() {

					classifier, err := NewClassifier("class1", "", "class2")
					Expect(err).Should(Equal(ErrClassNameEmpty))
					Expect(classifier).To(BeNil())
				})

			})
			Describe("positive test cases", func() {
				var classNames []string
				BeforeEach(func() {
					classNames = []string{"class1", "class2"}
				})
				It("should work for at least 2 distinct classes with no duplicates", func() {
					classifier, err := NewClassifier(classNames...)
					Expect(err).Should(BeNil())
					Expect(classifier).NotTo(BeNil())
				})
				It("should work for at least 2 distinct classes with duplicates present", func() {

					classifier, err := NewClassifier("class1", "class2", "class1")
					Expect(err).Should(BeNil())
					Expect(classifier).NotTo(BeNil())
				})
				It("should work for class names with special characters", func() {

					classifier, err := NewClassifier("class1", "~!@@#$%^&*()_+=-5gfdg", "class3")
					Expect(err).Should(BeNil())
					Expect(classifier).NotTo(BeNil())
				})
				It("should work for class names with unicode characters", func() {

					classifier, err := NewClassifier("class1", "क्लास २ ", "class1")
					Expect(err).Should(BeNil())
					Expect(classifier).NotTo(BeNil())
				})
			})

		})
	})

	Describe("train", func() {
		var classifier *Classifier
		Context("classifier not initialized", func() {
			It("should error", func() {
				classifier = &Classifier{}
				err := classifier.Train("some text")
				Expect(err).Should(Equal(ErrClassifierNotInitialized))
			})
		})
		Context("classifier initialized with 3 distinct classes", func() {
			BeforeEach(func() {
				classifier, _ = NewClassifier("class1", "class2", "class3")

			})
			It("should error for empty class name", func() {
				err := classifier.Train("", "some text")
				Expect(err).To(Equal(ErrClassNameEmpty))
			})
			It("should error for unknow class name", func() {
				err := classifier.Train("class4", "some text")
				Expect(err).To(Equal(ErrUnknownClassName("class4")))
			})
			It("should error for nil text", func() {
				err := classifier.Train("class3")
				Expect(err).To(Equal(ErrTrainingDataRequired))
			})
			It("should error for empty text", func() {
				err := classifier.Train("class3", "some data", "")
				Expect(err).To(Equal(ErrTrainingDataEmpty))
			})
			It("should error for empty text", func() {
				err := classifier.Train("class3", "")
				Expect(err).To(Equal(ErrTrainingDataEmpty))
			})
			It("should not error for valid text", func() {
				err := classifier.Train("class3", "some data")
				Expect(err).To(BeNil())
			})
			It("should not error for valid text with special characters", func() {
				err := classifier.Train("class3", "some data with ~!@#$%^&*()_+=-")
				Expect(err).To(BeNil())
			})
			It("should not error for valid text with unicode characters", func() {
				err := classifier.Train("class3", "some data with ₹ आणि अजुन शब्द")
				Expect(err).To(BeNil())
			})
		})

	})
	Describe("Finalized", func() {
		Context("Classifier not initialized", func() {
			It("should error", func() {
				classifier := &Classifier{}
				err := classifier.Finalize()
				Expect(err).Should(Equal(ErrClassifierNotInitialized))
			})
		})
		Context("Classifier not trained", func() {
			var classifier *Classifier
			BeforeEach(func() {
				classifier, _ = NewClassifier("class1", "class2", "class3")
			})
			It("should finalize without training", func() {
				err := classifier.Finalize()
				Expect(err).Should(BeNil())
			})
		})
		Context("Classifier trained with sample configuration", func() {
			var classifier *Classifier
			BeforeEach(func() {
				classifier, _ = NewClassifier("class1", "class2", "class3")
				classifier.Train("class1", "Credit Interest", "Interest Credited", "Realized Interest on xxxx")
				classifier.Train("class2", "Coffee", "StormBukcs", "DoMcnalds xxxx")
				classifier.Train("class3", "Rent", "House Rentxxxxx", "Repairwork")
			})
			It("should finalize with training", func() {
				err := classifier.Finalize()
				Expect(err).Should(BeNil())
			})
		})
	})
	Describe("Classify", func() {
		Context("Classifier not initialized", func() {
			It("should error", func() {
				classifier := &Classifier{}
				_, err := classifier.Classify("some text")
				Expect(err).Should(Equal(ErrClassifierNotInitialized))
			})
		})
		Context("Classifier not trained", func() {
			var classifier *Classifier
			BeforeEach(func() {
				classifier, _ = NewClassifier("class1", "class2", "class3")
			})
			It("should classify without training", func() {
				class, err := classifier.Classify("some text")
				Expect(err).Should(BeNil())
				Expect(class).ShouldNot(BeNil())
			})
		})
		Context("Classifier not finalized", func() {
			var classifier *Classifier
			BeforeEach(func() {
				classifier, _ = NewClassifier("class1", "class2", "class3")
				classifier.Train("class1", "Credit Interest", "Interest Credited", "Realized Interest on xxxx")
				classifier.Train("class2", "Coffee", "StormBukcs", "DoMcnalds xxxx")
				classifier.Train("class3", "Rent", "House Rentxxxxx", "Repairwork")
			})
			It("should classify with training not finalized", func() {
				class, err := classifier.Classify("some text")
				Expect(err).Should(BeNil())
				Expect(class).ShouldNot(BeNil())
			})
			It("should allow additional training after classification", func() {
				class, err := classifier.Classify("some text")
				Expect(err).Should(BeNil())
				Expect(class).ShouldNot(BeNil())
				err = classifier.Train("class1", "Interest payed")
				Expect(err).Should(BeNil())
				class, err = classifier.Classify("some more text")
				Expect(err).Should(BeNil())
				Expect(class).ShouldNot(BeNil())
			})
		})
		Context("Classifier trained and finalized with sample configuration", func() {
			var classifier *Classifier
			BeforeEach(func() {
				classifier, _ = NewClassifier("class2", "class1", "class3")
				classifier.Train("class1", "Credit Interest", "Interest Credited", "Realized Interest on XXXX")
				classifier.Train("class2", "Coffee", "StormBukcs", "DoMcnalds XXXX")
				classifier.Train("class3", "Rent payed", "House Rent XXXXXX", "payment of Rent")
				classifier.Finalize()
			})
			It("should classify seen text correctly", func() {
				class, err := classifier.Classify("StormBukcs")
				Expect(err).Should(BeNil())
				Expect(class).Should(Equal("class2"))
			})
			It("should classify generalized text correctly", func() {
				class, err := classifier.Classify("House Rent XXXXXX")
				Expect(err).Should(BeNil())
				Expect(class).Should(Equal("class3"))
			})
			It("for unknown text should return empty", func() {
				class, err := classifier.Classify("papaya")
				Expect(err).Should(BeNil())
				Expect(class).Should(Equal(""))
			})
		})
	})
})
