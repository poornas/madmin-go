//
// MinIO Object Storage (c) 2022 MinIO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

//go:build (linux && 386) || (linux && amd64) || (linux && arm64) || (linux && mips64) || (linux && mips)
// +build linux,386 linux,amd64 linux,arm64 linux,mips64 linux,mips

package kernel

func utsnameStr(in []int8) string {
	out := make([]byte, 0, len(in))
	for i := 0; i < len(in); i++ {
		if in[i] == 0x00 {
			break
		}
		out = append(out, byte(in[i]))
	}
	return string(out)
}
