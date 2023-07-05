/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metrics

import "dubbo.apache.org/dubbo-go/v3/common"

const (
	TagIp                    = "ip"
	TagPid                   = "pid"
	TagHostname              = "hostname"
	TagApplicationName       = "application.Name"
	TagApplicationModule     = "application.module.id"
	TagInterfaceKey          = "interface"
	TagMethodKey             = "method"
	TagGroupKey              = "group"
	TagVersionKey            = "Version"
	TagApplicationVersionKey = "application.Version"
	TagKeyKey                = "key"
	TagConfigCenter          = "config.center"
	TagChangeType            = "change.type"
	TagThreadName            = "thread.pool.Name"
	TagGitCommitId           = "git.commit.id"
)

type MetricKey struct {
	Name string
	Desc string
}

func NewMetricKey(name string, desc string) *MetricKey {
	return &MetricKey{Name: name, Desc: desc}
}

type MetricLevel interface {
	Tags() map[string]string
}

type ApplicationMetricLevel struct {
	Name        string
	Version     string
	GitCommitId string
}

func NewApplicationMetric(name string, version string, gitCommitId string) *ApplicationMetricLevel {
	return &ApplicationMetricLevel{
		Name:        name,
		Version:     version,
		GitCommitId: gitCommitId,
	}
}

func (m *ApplicationMetricLevel) Tags() map[string]string {
	return map[string]string{
		TagIp:                    common.GetLocalIp(),
		TagHostname:              common.GetLocalHostName(),
		TagApplicationName:       m.Name,
		TagApplicationVersionKey: m.Version,
		TagGitCommitId:           m.GitCommitId,
	}
}

type ServiceMetricLevel struct {
	*ApplicationMetricLevel
	InterfaceName string
}

func NewServiceMetric(name string, version string, gitCommitId string, interfaceName string) *ServiceMetricLevel {
	return &ServiceMetricLevel{
		ApplicationMetricLevel: NewApplicationMetric(name, version, gitCommitId),
		InterfaceName:          interfaceName,
	}
}

func (m *ServiceMetricLevel) Tags() map[string]string {
	tags := m.ApplicationMetricLevel.Tags()
	tags[TagInterfaceKey] = m.InterfaceName
	return tags
}

type MethodMetricLevel struct {
	*ServiceMetricLevel
	Method  string
	Group   string
	Version string
}

func (m *MethodMetricLevel) Tags() map[string]string {
	tags := m.ServiceMetricLevel.Tags()
	tags[TagMethodKey] = m.Method
	tags[TagGroupKey] = m.Group
	tags[TagVersionKey] = m.Version
	return tags
}

func NewMethodMetric(name string, version string, gitCommitId string, interfaceName string, method string, group string, methodVersion string) *MethodMetricLevel {
	return &MethodMetricLevel{
		ServiceMetricLevel: NewServiceMetric(name, version, gitCommitId, interfaceName),
		Method:             method,
		Group:              group,
		Version:            methodVersion,
	}
}
