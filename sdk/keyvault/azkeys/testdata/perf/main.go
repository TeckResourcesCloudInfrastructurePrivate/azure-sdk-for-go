// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import "github.com/Azure/azure-sdk-for-go/sdk/internal/perf"

func main() {
	perf.Run(map[string]perf.PerfMethods{
		"GetKeyTest":  {Register: nil, New: newGetKeyTest},
		"DecryptTest": {Register: nil, New: newDecryptTest},
		"SignTest":    {Register: nil, New: newSignTest},
		"UnwrapTest":  {Register: nil, New: newUnwrapTest},
	})
}
