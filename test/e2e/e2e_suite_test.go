package e2e_test

import (
	"flag"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/klog"
	"kubesphere.io/kubesphere/pkg/apis/network/v1alpha1"
	"kubesphere.io/kubesphere/pkg/test"
)

var ctx *test.TestCtx

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Networking E2e Suite")
}

var _ = BeforeSuite(func() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "false")
	flag.Set("v", "4")
	flag.Parse()
	klog.SetOutput(GinkgoWriter)

	ctx = test.NewTestCtx(nil, os.Getenv("TEST_NAMESPACE"))
	Expect(ctx.Setup(os.Getenv("YAML_PATH"), "", v1alpha1.AddToScheme)).ShouldNot(HaveOccurred())
	deployName := os.Getenv("DEPLOY_NAME")
	Expect(test.WaitForController(ctx.Client, ctx.Namespace, deployName, 1, time.Second*5, time.Minute)).ShouldNot(HaveOccurred(), "Controlller failed to start")
	klog.Infoln("Controller is up, begin to test ")
})

var _ = AfterSuite(func() {
	ctx.Cleanup(nil)
})
