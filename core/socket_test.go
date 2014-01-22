package core_test

import (
	"../util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	keigo "github.com/taichi/keigo/core"
)

var _ = Describe("Socket", func() {
	Context("when connect to Keiko succesfully", func() {
		var (
			config *keigo.KeigoConfig
			sock   keigo.Socket
			err    error
		)

		BeforeEach(util.ToOnceFn(func() {
			config = LoadTestConfig()
		}))

		BeforeEach(func() {
			Expect(err).NotTo(HaveOccurred())
			sock, err = keigo.NewSocket(config)
		})
		AfterEach(func() {
			err = sock.Close()
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have no error", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when connect to Keiko failure", func() {
		var config *keigo.KeigoConfig
		BeforeEach(func() {
			config = &keigo.KeigoConfig{
				Address:  "example.jp:3000",
				Terminal: '\r',
			}
		})
		It("should fail to connect", func() {
			_, e := keigo.NewSocket(config)
			Expect(e).To(HaveOccurred())
		})
	})
})
