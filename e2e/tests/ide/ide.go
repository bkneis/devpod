package ide

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/loft-sh/devpod/e2e/framework"
	"github.com/onsi/ginkgo/v2"
)

var _ = DevPodDescribe("devpod ide test suite", func() {
	ginkgo.Context("testing ides", ginkgo.Label("ide"), ginkgo.Ordered, func() {
		var initialDir string

		ginkgo.BeforeEach(func() {
			var err error
			initialDir, err = os.Getwd()
			framework.ExpectNoError(err)
		})

		// ginkgo.It("start ides", func() {
		// 	ctx := context.Background()

		// 	f := framework.NewDefaultFramework(initialDir + "/bin")
		// 	tempDir, err := framework.CopyToTempDir("tests/ide/testdata")
		// 	framework.ExpectNoError(err)
		// 	ginkgo.DeferCleanup(framework.CleanupTempDir, initialDir, tempDir)

		// 	_ = f.DevPodProviderDelete(ctx, "docker")
		// 	err = f.DevPodProviderAdd(ctx, "docker")
		// 	framework.ExpectNoError(err)
		// 	err = f.DevPodProviderUse(context.Background(), "docker")
		// 	framework.ExpectNoError(err)

		// 	ginkgo.DeferCleanup(f.DevPodWorkspaceDelete, context.Background(), tempDir)

		// 	err = f.DevPodUpWithIDE(ctx, tempDir, "--open-ide=false", "--ide=vscode")
		// 	framework.ExpectNoError(err)

		// 	err = f.DevPodUpWithIDE(ctx, tempDir, "--open-ide=false", "--ide=openvscode")
		// 	framework.ExpectNoError(err)

		// 	err = f.DevPodUpWithIDE(ctx, tempDir, "--open-ide=false", "--ide=jupyternotebook")
		// 	framework.ExpectNoError(err)

		// 	err = f.DevPodUpWithIDE(ctx, tempDir, "--open-ide=false", "--ide=fleet")
		// 	framework.ExpectNoError(err)

		// 	// check if ssh works
		// 	err = f.DevPodSSHEchoTestString(ctx, tempDir)
		// 	framework.ExpectNoError(err)

		// 	// Set up port forwarding since IDE was not opened
		// 	go func() {
		// 		_ = f.DevpodPortTest(ctx, "10700", tempDir)
		// 		// _ = f.DevpodPortTest(ctx, "10800", tempDir)
		// 	}()
		// 	time.Sleep(time.Second)

		// 	// Curl the ports and ensure they respond with HTML
		// 	port := 10700
		// 	address := fmt.Sprintf("http://localhost:%d", port)
		// 	res, err := http.DefaultClient.Get(address)
		// 	framework.ExpectNoError(err)
		// 	framework.ExpectEqual(res.Header.Get("content-type"), "text/html; charset=UTF-8")

		// 	// Set up port forwarding since IDE was not opened
		// 	go func() {
		// 		_ = f.DevpodPortTest(ctx, "10800", tempDir)
		// 		// _ = f.DevpodPortTest(ctx, "10800", tempDir)
		// 	}()
		// 	time.Sleep(time.Second)

		// 	// Curl the ports and ensure they respond with HTML
		// 	port = 10800
		// 	address = fmt.Sprintf("http://localhost:%d", port)
		// 	res, err = http.DefaultClient.Get(address)
		// 	framework.ExpectNoError(err)
		// 	framework.ExpectEqual(res.Header.Get("content-type"), "text/html")

		// 	// TODO: test jetbrains ides
		// })

		ginkgo.It("ensures open vscode is available locally when started using the platform", func() {
			ctx := context.Background()

			f := framework.NewDefaultFramework(initialDir + "/bin")
			// tempDir, err := framework.CopyToTempDir("tests/ide/testdata")
			// framework.ExpectNoError(err)
			// ginkgo.DeferCleanup(framework.CleanupTempDir, initialDir, tempDir)

			// _ = f.DevPodProviderDelete(ctx, "docker")
			// err = f.DevPodProviderAdd(ctx, "docker")
			// framework.ExpectNoError(err)
			// err = f.DevPodProviderUse(context.Background(), "docker")
			// framework.ExpectNoError(err)

			err := f.DevPodProLogin(ctx)
			framework.ExpectNoError(err)
			go func() {
				err := f.DevPodDaemonStart(ctx)
				framework.ExpectNoError(err)
			}()
			time.Sleep(2 * time.Second)

			ginkgo.DeferCleanup(f.DevPodProDelete, context.Background())

			ginkgo.DeferCleanup(f.DevPodWorkspaceDelete, context.Background(), "testprovscode")

			err = f.DevPodProUpWithIDE(ctx, "testprovscode", "git", "https://github.com/loft-sh/devpod-example-go", "--open-ide=false", "--ide=openvscode")
			framework.ExpectNoError(err)

			// check if ssh works
			err = f.DevPodSSHEchoTestString(ctx, "testprovscode")
			framework.ExpectNoError(err)

			// Set up port forward since IDE was not opened
			go func() {
				_ = f.DevpodPortTest(ctx, "10800", "testprovscode")
			}()
			time.Sleep(2 * time.Second)

			// Curl the ports and ensure they respond with HTML
			port := 10800
			address := fmt.Sprintf("http://localhost:%d", port)
			res, err := http.DefaultClient.Get(address)
			framework.ExpectNoError(err)
			framework.ExpectEqual(res.Header.Get("content-type"), "text/html")

			// err = f.DevPodProUpWithIDE(ctx, "testprojupyternotebook", "image", "python:3.13.3-bookworm", "--open-ide=false", "--ide=jupyternotebook")
			// framework.ExpectNoError(err)

			// // check if ssh works
			// err = f.DevPodSSHEchoTestString(ctx, "testprojupyternotebook")
			// framework.ExpectNoError(err)

			// // Set up port forward since IDE was not opened
			// go func() {
			// 	_ = f.DevpodPortTest(ctx, "10700", "testprojupyternotebook")
			// }()
			// time.Sleep(2 * time.Second)

			// // Curl the ports and ensure they respond with HTML
			// port = 10700
			// address = fmt.Sprintf("http://localhost:%d", port)
			// res, err = http.DefaultClient.Get(address)
			// framework.ExpectNoError(err)
			// framework.ExpectEqual(res.Header.Get("content-type"), "text/html; charset=UTF-8")
		})

		ginkgo.It("ensures open jupyter notebook is available locally when started using the platform", func() {
			ctx := context.Background()

			f := framework.NewDefaultFramework(initialDir + "/bin")
			err := f.DevPodProLogin(ctx)
			framework.ExpectNoError(err)
			go func() {
				err := f.DevPodDaemonStart(ctx)
				framework.ExpectNoError(err)
			}()
			time.Sleep(2 * time.Second)

			ginkgo.DeferCleanup(f.DevPodProDelete, context.Background())

			ginkgo.DeferCleanup(f.DevPodWorkspaceDelete, context.Background(), "testrojupyternotebook")

			err = f.DevPodProUpWithIDE(ctx, "testprojupyternotebook", "image", "python:3.13.3-bookworm", "--open-ide=false", "--ide=jupyternotebook")
			framework.ExpectNoError(err)

			// check if ssh works
			err = f.DevPodSSHEchoTestString(ctx, "testprojupyternotebook")
			framework.ExpectNoError(err)

			// Set up port forward since IDE was not opened
			go func() {
				_ = f.DevpodPortTest(ctx, "10700", "testprojupyternotebook")
			}()
			time.Sleep(2 * time.Second)

			// Curl the ports and ensure they respond with HTML
			port := 10700
			address := fmt.Sprintf("http://localhost:%d", port)
			res, err := http.DefaultClient.Get(address)
			framework.ExpectNoError(err)
			framework.ExpectEqual(res.Header.Get("content-type"), "text/html; charset=UTF-8")
		})

		// ginkgo.It("ensures jupyter notebook is available locally", func() {
		// 	ctx := context.Background()

		// 	f := framework.NewDefaultFramework(initialDir + "/bin")
		// 	tempDir, err := framework.CopyToTempDir("tests/ide/testdata")
		// 	framework.ExpectNoError(err)
		// 	ginkgo.DeferCleanup(framework.CleanupTempDir, initialDir, tempDir)

		// 	_ = f.DevPodProviderDelete(ctx, "docker")
		// 	err = f.DevPodProviderAdd(ctx, "docker")
		// 	framework.ExpectNoError(err)
		// 	err = f.DevPodProviderUse(context.Background(), "docker")
		// 	framework.ExpectNoError(err)

		// 	ginkgo.DeferCleanup(f.DevPodWorkspaceDelete, context.Background(), tempDir)

		// 	err = f.DevPodUpWithIDE(ctx, tempDir, "--open-ide=false", "--ide=jupyternotebook")
		// 	framework.ExpectNoError(err)

		// 	// check if ssh works
		// 	err = f.DevPodSSHEchoTestString(ctx, tempDir)
		// 	framework.ExpectNoError(err)

		// 	// Set up port forward since IDE was not opened
		// 	go func() {
		// 		_ = f.DevpodPortTest(ctx, "10700", tempDir)
		// 	}()
		// 	time.Sleep(2 * time.Second)

		// 	// Curl the ports and ensure they respond with HTML
		// 	port := 10700
		// 	address := fmt.Sprintf("http://localhost:%d", port)
		// 	res, err := http.DefaultClient.Get(address)
		// 	framework.ExpectNoError(err)
		// 	framework.ExpectEqual(res.Header.Get("content-type"), "text/html; charset=UTF-8")
		// })
	})
})
