// Copyright 2020 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// To run locally you need to export the env var with the image
// $ export MEMCACHED_IMAGE=memcached:1.4.36-alpine

package e2e_go_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo" //nolint:golint
	. "github.com/onsi/gomega" //nolint:golint
)

var _ = Describe("Running Go projects", func() {
	Context("built with operator-sdk", func() {

		BeforeEach(func() {
			By("installing CRD's")
			err := tc.Make("install")
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			By("uninstalling CRD's")
			err := tc.Make("uninstall")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should run correctly locally", func() {
			By("running the project")
			cmd := exec.Command("make", "run")
			err := cmd.Start()
			Expect(err).NotTo(HaveOccurred())

			By("killing the project")
			err = cmd.Process.Kill()
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
