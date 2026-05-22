// Copyright 2025.
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

//go:build windows

package services

import "syscall"

// setDetach is a no-op on Windows because syscall.SysProcAttr does not have Setsid.
func setDetach(attr *syscall.SysProcAttr) {
	// No detachment tweak needed; Windows uses different process group semantics.
}
