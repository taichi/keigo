package core_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CKDI", func() {
	It("should work normally", func() {
		cmd := &CKDI{States{}}
		Expect(cmd.String()).To(Equal("CKDI"))
	})
	It("should work with states", func() {
		cmd := &CKDI{States{On, Disable, Keep}}
		Expect(cmd.String()).To(Equal("CKDI EDxx"))
	})
})
