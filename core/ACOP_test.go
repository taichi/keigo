package core_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ACOP", func() {
	It("should work normally", func() {
		cmd := &ACOP{DirectOutput, 2, 1, States{}}
		Expect(cmd.String()).To(Equal("ACOP -u 2 -w 2 -t 1"))
	})
	It("should work with states", func() {
		cmd := &ACOP{LampOrBuzzer, 2, 1, States{On, On, Blink}}
		Expect(cmd.String()).To(Equal("ACOP -u 1 112xxxxx -w 2 -t 1"))
	})
})
