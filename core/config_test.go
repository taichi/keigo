package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/keigo/core"
)

func LoadTestConfig() *KeigoConfig {
	config, err := LoadConfig("config.toml")
	if err != nil {
		panic(err)
	}
	return config
}

var _ = Describe("Config", func() {
	Context("When load Config succesfully", func() {
		var (
			config *KeigoConfig
			err    error
		)
		BeforeEach(func() {
			config, err = LoadConfig("config.toml")
		})
		It("should not have error", func() {
			Expect(err).NotTo(HaveOccurred())
		})
		It("should have values", func() {
			Expect(config.Address).To(Equal("keiko:60000"))
			Expect(config.Terminal).To(BeEquivalentTo('\r'))
		})
	})
})
