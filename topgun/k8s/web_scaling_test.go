package k8s_test

import (
	"strconv"
	"time"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scaling web instances", func() {

	BeforeEach(func() {
		setReleaseNameAndNamespace("swi")
	})

	AfterEach(func() {
		cleanup(releaseName, namespace, nil)
	})

	It("succeeds", func() {
		successfullyDeploysConcourse(1, 1)
		successfullyDeploysConcourse(0, 1)
		successfullyDeploysConcourse(2, 1)
	})
})

func successfullyDeploysConcourse(webReplicas, workerReplicas int) {
	deployConcourseChart(releaseName,
		"--set=web.replicas="+strconv.Itoa(webReplicas),
		"--set=worker.replicas="+strconv.Itoa(workerReplicas),
	)

	waitAllPodsInNamespaceToBeReady(namespace)

	By("Creating the web proxy")
	proxySession, atcEndpoint := startPortForwarding(namespace, "service/"+releaseName+"-web", "8080")
	defer proxySession.Interrupt()

	By("Logging in")
	fly.Login("test", "test", atcEndpoint)

	By("waiting for a running worker")
	Eventually(func() []Worker {
		return getRunningWorkers(fly.GetWorkers())
	}, 5*time.Minute, 60*time.Second).
		Should(HaveLen(workerReplicas))
}
