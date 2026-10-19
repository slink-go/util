// SPDX-License-Identifier: AGPL-3.0-only
// Copyright (C) 2024-2026 Mikhail Kantur

package date

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDateMarshalling(t *testing.T) {
	var v = struct {
		Title string `json:"title"`
		Date  Date   `json:"date"`
	}{
		Title: "test",
		Date:  Now(),
	}
	data, err := json.Marshal(v)
	fmt.Println(string(data), err)
}
