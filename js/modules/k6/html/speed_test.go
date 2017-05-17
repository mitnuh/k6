/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package html

import (
	// 	"fmt"
	"context"
	"errors"
	"testing"

	"github.com/dop251/goja"
	"github.com/loadimpact/k6/js/common"
)

const benchmarkElemHTML = `
<html>
<head>
	<title>This is the title</title>
</head>
<body>
	<p>Lorem</p> <p>ipsum</p> <p>dolor</p> <p>sit</p> <p>amet,</p> <p>consectetur</p> <p>adipiscing</p> <p>elit.</p>
	innerfirst<h2 id="h2_elem" class="class2">Nullam id nisi eget ex pharetra imperdiet.</h2>
	<span><b>test content</b></span>innerlast
</body>
`

func buildElemBenchmark(num int, rt *goja.Runtime, prg *goja.Program, b *testing.B) {
	for i := 0; i < num; i++ {
		if _, err := rt.RunProgram(prg); err != nil {
			panic(errors.New("Unable to create element"))
		}
	}
}

func BenchmarkBuildSpeed(b *testing.B) {
	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper{})

	ctx := common.WithRuntime(context.Background(), rt)
	rt.Set("src", benchmarkElemHTML)
	rt.Set("html", common.Bind(rt, &HTML{}, &ctx))
	compileProtoElem()

	if _, err := common.RunString(rt, `let doc = html.parseHTML(src)`); err != nil {
		return
	}

	if _, err := common.RunString(rt, `let body = doc.find("body")`); err != nil {
		return
	}

	prg := common.MustCompile("GetElem", `body.get(0)`, true)
	b.Run("BenchmarkElement100", func(b *testing.B) { buildElemBenchmark(100, rt, prg, b) })
	b.Run("BenchmarkElement500", func(b *testing.B) { buildElemBenchmark(100, rt, prg, b) })
	b.Run("BenchmarkElement1000", func(b *testing.B) { buildElemBenchmark(1000, rt, prg, b) })
	b.Run("BenchmarkElement5000", func(b *testing.B) { buildElemBenchmark(5000, rt, prg, b) })
	b.Run("BenchmarkElement10000", func(b *testing.B) { buildElemBenchmark(10000, rt, prg, b) })
	// b.Run("BenchmarkElement100000", func(b *testing.B) { buildElemBenchmark(100000, rt, prg, b) })
}
