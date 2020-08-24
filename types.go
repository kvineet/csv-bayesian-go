package main

import (
	"fmt"
	"sort"

	"github.com/navossoc/bayesian"
)

var (
	//ErrNotEnoughClassesDefined indicates error in classifier initialization where no classes are defined
	ErrNotEnoughClassesDefined = fmt.Errorf("At least two unique classes sgould be defined")
	//ErrClassNameEmpty is error when class name provided is empty
	ErrClassNameEmpty = fmt.Errorf("Class name can not be empty")
	//ErrClassifierNotInitialized is error when classifier is used before initialization
	ErrClassifierNotInitialized = fmt.Errorf("Classifier is not initialized")
	//ErrUnknownClassName indicates classifier sees a new class while training
	ErrUnknownClassName = func(name string) error { return fmt.Errorf("Class %s is unknown", name) }
	//ErrTrainingDataRequired indicates error where classfier was trained wihtout data
	ErrTrainingDataRequired = fmt.Errorf("Training document can not be empty")
	//ErrTrainingDataEmpty indicates error when training data is empty
	ErrTrainingDataEmpty = fmt.Errorf("Training data can not be empty")
)

// Classifier is a wrapper for bayesian classifier implementation
type Classifier struct {
	clf     *bayesian.Classifier
	classes map[string]bayesian.Class
}

//NewClassifier instantiates new classifier
func NewClassifier(classNames ...string) (*Classifier, error) {
	if len(classNames) < 2 {
		return nil, ErrNotEnoughClassesDefined
	}
	classMap := map[string]bayesian.Class{}
	classes := []bayesian.Class{}

	for _, v := range classNames {
		if v == "" {
			return nil, ErrClassNameEmpty
		}
		if _, val := classMap[v]; val {
			continue
		}
		class := bayesian.Class(v)
		classes = append(classes, class)
		classMap[v] = class
	}
	if len(classes) < 2 {
		return nil, ErrNotEnoughClassesDefined
	}
	c := bayesian.NewClassifier(classes...)
	return &Classifier{clf: c, classes: classMap}, nil
}

//Train function trains the classifier with provided data
func (in *Classifier) Train(className string, data ...string) error {
	if in.clf == nil {
		return ErrClassifierNotInitialized
	}
	if className == "" {
		return ErrClassNameEmpty
	}
	if len(data) < 1 {
		return ErrTrainingDataRequired
	}
	class, ok := in.classes[className]
	if !ok {
		return ErrUnknownClassName(className)
	}
	for _, doc := range data {
		if doc == "" {
			return ErrTrainingDataEmpty
		}
		in.clf.Learn([]string{doc}, class)
	}
	return nil
}

//Finalize finalizes training
func (in *Classifier) Finalize() error {
	if in.clf == nil {
		return ErrClassifierNotInitialized
	}
	in.clf.ConvertTermsFreqToTfIdf()
	return nil
}

// Classify classified given text in learned categories
func (in *Classifier) Classify(text string) (string, error) {
	if in.clf == nil {
		return "", ErrClassifierNotInitialized
	}
	scores, likely, strict, err := in.clf.SafeProbScores([]string{text})
	if err != nil {
		return "", nil
	}
	if !strict {
		sort.Slice(scores, func(i, j int) bool {
			return scores[i] > scores[j]
		})
		if FloatEquals(scores[0], scores[1]) {
			return "", nil
		}
	}
	class := in.clf.Classes[likely]
	return string(class), nil
}
