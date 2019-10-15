package k8s_test

import (
	"time"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("TSA Service Node Port", func() {

	var (
		proxySession *gexec.Session
	)

	JustBeforeEach(func() {
		setReleaseNameAndNamespace("tnp")

		deployConcourseChart(releaseName,
			"--set=web.service.type=NodePort",
		)
	})

	It("deployment succeeds", func() {
		var atcEndpoint string

		proxySession, atcEndpoint = startPortForwarding(namespace, "service/"+releaseName+"-web", "8080")

		fly.Login("test", "test", atcEndpoint)
		Eventually(func() []Worker {
			return getRunningWorkers(fly.GetWorkers())
		}, 5*time.Minute, 60*time.Second).
			ShouldNot(HaveLen(0))
	})

	AfterEach(func() {
		cleanup(releaseName, namespace, proxySession)
	})

})
