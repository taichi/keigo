package core_test

import (
	"../util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	keigo "github.com/taichi/keigo/core"
	"time"
)

var _ = Describe("Session", func() {
	var (
		config  *keigo.KeigoConfig
		session keigo.Session
		err     error
	)

	BeforeEach(util.ToOnceFn(func() {
		config = LoadTestConfig()
	}))
	AfterEach(func() {
		session.Close()
	})

	connectFn := func(second time.Duration) func() {
		return func() {
			Expect(err).NotTo(HaveOccurred())
			config.Timeout = second * time.Second
			session, err = keigo.Connect(config)
			Expect(err).NotTo(HaveOccurred())
		}
	}

	Context("when connect to Keiko succesfully", func() {
		BeforeEach(connectFn(60))
		Describe("CKST", func() {
			It("should get SNMP trap status", func() {
				v, e := session.Execute(keigo.CKST)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[ED]{20}"))
			})
		})
		Describe("RDCD", func() {
			It("should get Effective Date", func() {
				v, e := session.Execute(keigo.RDCD)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("([0-9]{4}/0[1-9]|1[012]/([012][0-9]|3[01]))|Not registered"))
			})
		})
		Describe("RDCN", func() {
			It("should get Contract Number", func() {
				v, e := session.Execute(keigo.RDCN)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[0-9]+|Not registered"))
			})
		})
		Describe("RDMN", func() {
			It("should get Model Name", func() {
				v, e := session.Execute(keigo.RDMN)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("DN-[0-9]{4}[A-Z]{2}"))
			})
		})
		Describe("RDPD", func() {
			It("should get Production Date", func() {
				v, e := session.Execute(keigo.RDPD)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[0-9]{2}(0[1-9]|1[012])"))
			})
		})
		Describe("RDSN", func() {
			It("should get Serial Number", func() {
				v, e := session.Execute(keigo.RDSN)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[0-9]+"))
			})
		})
		Describe("ROPS", func() {
			It("should get Direct input state", func() {
				v, e := session.Execute(keigo.ROPS)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[01]{4}"))
			})
		})
		Describe("UTID", func() {
			It("should get unit id", func() {
				v, e := session.Execute(keigo.UTID)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("[0-9]+"))
			})
		})
		Describe("VERN", func() {
			It("should get current version", func() {
				v, e := session.Execute(keigo.VERN)
				Expect(e).NotTo(HaveOccurred())
				Expect(v).To(MatchRegexp("14\\..*"))
			})
		})
	})
	Describe("execute with Reconnect", func() {
		request := func() {
			time.Sleep(config.Timeout + 2*time.Second)
			id, err := session.Execute(keigo.UTID)
			Expect(err).NotTo(HaveOccurred())
			Expect(id).NotTo(BeEmpty())
		}
		Context("when disconnect by Client", func() {
			BeforeEach(connectFn(3))
			It("should reconnect to Keiko", request)
		})
		Context("when disconnect by Server", func() {
			BeforeEach(connectFn(60))
			It("should reconnect to Keiko", request)
		})
	})
})
