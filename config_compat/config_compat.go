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

package config_compat

import "dubbo.apache.org/dubbo-go/v3/common"

// Load is used to resolve circular dependencies temporarily.
// source : compat.go func compatLoad(opts ...interface{}) error
var Load func(opts ...interface{}) error

// todo compat LoaderConfOption

// SetProviderService is used to resolve circular dependencies temporarily.
// source : compat.go func compatSetProviderService(srv common.RPCService)
var SetProviderService func(service common.RPCService)

// SetConsumerService is used to resolve circular dependencies temporarily.
// source : compat.go func compatSetProviderService(srv common.RPCService)
var SetConsumerService func(service common.RPCService)
