package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Animals river testing", func() {

	Describe("Testing position change", func() {

		var entitiesTest []Entity

		BeforeEach(func() {
			entitiesTest = entities
		})

		It("should return other position", func() {
			currPos := entitiesTest[0].Position
			moveEntity(&entitiesTest[0])

			Expect(currPos).To(Not(Equal(entities[0].Position)))
		})

		It("should return other position", func() {
			currPos := entitiesTest[1].Position
			moveEntity(&entitiesTest[1])

			Expect(currPos).To(Not(Equal(entities[1].Position)))
		})
	})

	Describe("Testing isValidState()", func() {

		It("should return false when state exist in antisates", func() {
			state := []Entity{
				{Goat, Right},
				{Wolf, Right},
				{Carrot, Left},
				{Cowboy, Left},
			}

			Expect(isValidState(state)).To(Equal(false))
		})

		It("should return true when state not exist in antisates", func() {
			state := []Entity{
				{Goat, Right},
				{Wolf, Left},
				{Carrot, Left},
				{Cowboy, Left},
			}

			Expect(isValidState(state)).To(Equal(true))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aspire Suite")
}
