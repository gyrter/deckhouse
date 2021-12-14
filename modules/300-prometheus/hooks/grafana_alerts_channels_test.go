/*
Copyright 2021 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = FDescribe("Prometheus hooks :: grafana notification channels ::", func() {
	const (
		testAlertsChannelYAML = `
apiVersion: deckhouse.io/v1alpha1
kind: GrafanaAlertsChannel
metadata:
  name: test
spec:
  type: prometheus-alertmanager
  alertManager:
    address: "http://test-alert-manager-url"
    auth:
      basic:
        username: user
        password: password
`
		testAlertsChanelUpdatedYAML = `
apiVersion: deckhouse.io/v1alpha1
kind: GrafanaAlertsChannel
metadata:
  name: test
spec:
  type: prometheus-alertmanager
  alertManager:
    address: "https://new-test-url"
    auth:
      basic:
        username: user
        password: new-password
`
		testAlertsChannelWithoutAuthYAML = `
apiVersion: deckhouse.io/v1alpha1
kind: GrafanaAlertsChannel
metadata:
  name: another
spec:
  type: prometheus-alertmanager
  alertManager:
    address: "https://another-url"
`
	)

	var (
		testAlertsChannel = GrafanaAlertsChannel{
			OrgID:                 1,
			Type:                  alertManagerGrafanaAlertChannelType,
			Name:                  "test",
			UID:                   "test",
			IsDefault:             false,
			DisableResolveMessage: false,
			SendReminder:          false,
			Frequency:             time.Duration(0),
			Settings: map[string]interface{}{
				"url":           "http://test-alert-manager-url",
				"basicAuthUser": "user",
			},
			SecureSettings: map[string]interface{}{
				"basicAuthPassword": "password",
			},
		}

		testAlertsChanelUpdated = GrafanaAlertsChannel{
			OrgID:                 1,
			Type:                  alertManagerGrafanaAlertChannelType,
			Name:                  "test",
			UID:                   "test",
			IsDefault:             false,
			DisableResolveMessage: false,
			SendReminder:          false,
			Frequency:             time.Duration(0),
			Settings: map[string]interface{}{
				"url":           "https://new-test-url",
				"basicAuthUser": "user",
			},
			SecureSettings: map[string]interface{}{
				"basicAuthPassword": "new-password",
			},
		}

		madisonAlertsChannel = GrafanaAlertsChannel{
			OrgID:                 1,
			Type:                  alertManagerGrafanaAlertChannelType,
			Name:                  madisonAlertChannelName,
			UID:                   madisonAlertChannelName,
			IsDefault:             false,
			DisableResolveMessage: false,
			SendReminder:          false,
			Frequency:             time.Duration(0),
			Settings: map[string]interface{}{
				"url": "http://madison-proxy.d8-monitoring.svc.cluster.my:8080",
			},
			SecureSettings: make(map[string]interface{}),
		}

		testAlertsChannelWithoutAuth = GrafanaAlertsChannel{
			OrgID:                 1,
			Type:                  alertManagerGrafanaAlertChannelType,
			Name:                  "another",
			UID:                   "another",
			IsDefault:             false,
			DisableResolveMessage: false,
			SendReminder:          false,
			Frequency:             time.Duration(0),
			Settings: map[string]interface{}{
				"url": "https://another-url",
			},
			SecureSettings: make(map[string]interface{}),
		}
	)

	f := HookExecutionConfigInit(`
{
  "global": {
    "enabledModules": [],
    "discovery":{
		"clusterDomain": "cluster.my"
    }
  },
  "prometheus":{
    "internal":{
      "grafana":{}
    }
 }
}`, ``)
	f.RegisterCRD("deckhouse.io", "v1alpha1", "GrafanaAlertsChannel", false)

	assertChannelsInValues := func(f *HookExecutionConfig, expectChannels []GrafanaAlertsChannel) {
		channels := f.ValuesGet("prometheus.internal.grafana.alertsChannels").Array()

		Expect(channels).To(HaveLen(len(expectChannels)))

		nameToChannel := make(map[string]GrafanaAlertsChannel)
		for _, c := range expectChannels {
			nameToChannel[c.UID] = c
		}

		for _, raw := range channels {
			c := GrafanaAlertsChannel{}
			err := json.Unmarshal([]byte(raw.Raw), &c)
			Expect(err).ToNot(HaveOccurred())

			expected, ok := nameToChannel[c.UID]

			Expect(ok).To(BeTrue())
			Expect(expected).To(Equal(c))
		}
	}

	Context("Empty cluster", func() {
		BeforeEach(func() {
			f.BindingContexts.Set(f.KubeStateSet(``))
			f.RunHook()
		})

		It("Does not set any channels in values", func() {
			Expect(f).To(ExecuteSuccessfully())

			assertChannelsInValues(f, make([]GrafanaAlertsChannel, 0))
		})

		Context("Add channel", func() {
			BeforeEach(func() {
				f.BindingContexts.Set(f.KubeStateSetAndWaitForBindingContexts(testAlertsChannelYAML, 1))
				f.RunHook()
			})

			It("Should store channel in values", func() {
				Expect(f).To(ExecuteSuccessfully())

				assertChannelsInValues(f, []GrafanaAlertsChannel{testAlertsChannel})
			})

			Context("Deleting channel", func() {
				BeforeEach(func() {
					f.BindingContexts.Set(f.KubeStateSetAndWaitForBindingContexts(``, 0))
					f.RunHook()
				})

				It("Should delete GrafanaAdditionalDatasource from values", func() {
					Expect(f).To(ExecuteSuccessfully())

					assertChannelsInValues(f, make([]GrafanaAlertsChannel, 0))
				})
			})

			Context("Updating channel", func() {
				BeforeEach(func() {
					f.BindingContexts.Set(f.KubeStateSetAndWaitForBindingContexts(testAlertsChanelUpdatedYAML, 1))
					f.RunHook()
				})

				It("Should update GrafanaAdditionalDatasource in values", func() {
					Expect(f).To(ExecuteSuccessfully())

					assertChannelsInValues(f, []GrafanaAlertsChannel{testAlertsChanelUpdated})
				})
			})
		})

		Context("Enable flant integration module", func() {
			BeforeEach(func() {
				f.ValuesSetFromYaml("global.enabledModules", []byte(`["flant-integration"]`))
				f.RunHook()
			})

			It("Should store madison alerts channel in values with url only", func() {
				Expect(f).To(ExecuteSuccessfully())

				assertChannelsInValues(f, []GrafanaAlertsChannel{madisonAlertsChannel})
			})

			Context("Disable flant integration module", func() {
				BeforeEach(func() {
					f.ValuesSetFromYaml("global.enabledModules", []byte(`[]`))
					f.RunHook()
				})

				It("Should remove madison alerts channel from values", func() {
					Expect(f).To(ExecuteSuccessfully())
					assertChannelsInValues(f, make([]GrafanaAlertsChannel, 0))
				})
			})
		})
	})

	Context("Alerts channels in cluster", func() {
		BeforeEach(func() {
			JoinKubeResourcesAndSet(f, testAlertsChannelYAML, testAlertsChannelWithoutAuthYAML)
			f.RunHook()
		})

		It("Should store all alerts channel into values", func() {
			Expect(f).To(ExecuteSuccessfully())

			assertChannelsInValues(f, []GrafanaAlertsChannel{testAlertsChannel, testAlertsChannelWithoutAuth})
		})
	})

	Context("Flant integration module is enabled", func() {
		BeforeEach(func() {
			f.ValuesSetFromYaml("global.enabledModules", []byte(`["flant-integration"]`))
			f.RunHook()
		})

		It("Should store madison alerts channel in values with url only", func() {
			Expect(f).To(ExecuteSuccessfully())

			assertChannelsInValues(f, []GrafanaAlertsChannel{madisonAlertsChannel})
		})
	})
})
