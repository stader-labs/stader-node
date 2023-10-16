/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.3.0]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package service

import (
	stdCf "github.com/stader-labs/stader-node/shared/services/config"
)

func setUIMonitoring(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	newSettings[keys.Nm_enable_metrics] = cfg.EnableMetrics.Value.(bool)

	exposeGuardianPort, ok := cfg.ExposeGuardianPort.Value.(bool)
	if ok {
		newSettings[keys.Nm_expose_guardian_port] = exposeGuardianPort
	} else {
		newSettings[keys.Nm_expose_guardian_port] = true
	}
	newSettings[keys.Nm_execution_client_metrics_port] = format(cfg.EcMetricsPort.Value)
	newSettings[keys.Nm_beacon_node_metrics_port] = format(cfg.BnMetricsPort.Value)
	newSettings[keys.Nm_validator_client_metrics_port] = format(cfg.VcMetricsPort.Value)
	newSettings[keys.Nm_node_metrics_port] = format(cfg.NodeMetricsPort.Value)
	newSettings[keys.Nm_exporter_metrics_port] = format(cfg.ExporterMetricsPort.Value)

	// Nm_guardian_oracle_port ??
	// Nm_allow_root_filesystem_access ??
	// Nm_enable_oracle_dao_metrics?

	//Grafana
	newSettings[keys.Nm_grafana_port] = format(cfg.Grafana.Port.Value)
	newSettings[keys.Nm_grafana_container_tag] = cfg.Grafana.ContainerTag.Value

	//Prometheus
	newSettings[keys.Nm_prometheus_port] = format(cfg.Prometheus.Port.Value)
	newSettings[keys.Nm_expose_prometheus_port] = cfg.Prometheus.OpenPort.Value.(bool)
	newSettings[keys.Nm_prometheus_container_tag] = cfg.Prometheus.ContainerTag.Value
	newSettings[keys.Nm_additional_prometheus_flags] = cfg.Prometheus.AdditionalFlags.Value

	newSettings[keys.Nm_exporter_container_tag] = cfg.Exporter.ContainerTag.Value
	newSettings[keys.Nm_additional_exporter_flags] = cfg.Exporter.AdditionalFlags.Value

	newSettings[keys.Nm_enable_beaconchain_node_metrics] = cfg.EnableBitflyNodeMetrics.Value.(bool)

	newSettings[keys.Nm_beaconchain_api_key] =
		cfg.BitflyNodeMetrics.Secret.Value
	newSettings[keys.Nm_beaconchain_node_metrics_machine_name] =
		cfg.BitflyNodeMetrics.MachineName.Value

	return nil
}

func updateMonitoring(cfg *stdCf.StaderConfig, newSettings map[string]interface{}) error {
	cfg.EnableMetrics.Value = newSettings[keys.Nm_enable_metrics]
	cfg.ExposeGuardianPort.Value = newSettings[keys.Nm_expose_guardian_port]
	if cfg.EnableMetrics.Value.(bool) == false {
		return nil
	}

	cfg.EcMetricsPort.Value = newSettings[keys.Nm_execution_client_metrics_port]
	cfg.BnMetricsPort.Value = newSettings[keys.Nm_beacon_node_metrics_port]
	cfg.VcMetricsPort.Value = newSettings[keys.Nm_validator_client_metrics_port]
	cfg.NodeMetricsPort.Value = newSettings[keys.Nm_node_metrics_port]
	cfg.ExporterMetricsPort.Value = newSettings[keys.Nm_exporter_metrics_port]

	// cfg.GuardianMetricsPort.Value =

	// Nm_guardian_oracle_port ??
	// Nm_allow_root_filesystem_access ??
	// Nm_enable_oracle_dao_metrics?

	//Grafana
	cfg.Grafana.Port.Value = newSettings[keys.Nm_grafana_port]
	cfg.Grafana.ContainerTag.Value = newSettings[keys.Nm_grafana_container_tag]

	//Prometheus
	cfg.Prometheus.Port.Value = newSettings[keys.Nm_prometheus_port]
	cfg.Prometheus.OpenPort.Value = newSettings[keys.Nm_expose_prometheus_port]
	cfg.Prometheus.ContainerTag.Value = newSettings[keys.Nm_prometheus_container_tag]
	cfg.Prometheus.AdditionalFlags.Value = newSettings[keys.Nm_additional_prometheus_flags]

	// Exporter
	cfg.Exporter.ContainerTag.Value = newSettings[keys.Nm_exporter_container_tag]
	cfg.Exporter.AdditionalFlags.Value = newSettings[keys.Nm_additional_exporter_flags]

	cfg.EnableBitflyNodeMetrics.Value =
		newSettings[keys.Nm_enable_beaconchain_node_metrics]

	cfg.BitflyNodeMetrics.Secret.Value =
		newSettings[keys.Nm_beaconchain_api_key]
	cfg.BitflyNodeMetrics.MachineName.Value =
		newSettings[keys.Nm_beaconchain_node_metrics_machine_name]
	return nil
}
