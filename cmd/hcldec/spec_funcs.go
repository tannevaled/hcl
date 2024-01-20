// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

var specFuncs = map[string]function.Function{
        // https://github.com/zclconf/go-cty/cty/function/stdlib/bool.go
	"not":             stdlib.NotFunc,
        "and":             stdlib.AndFunc,
	"or":              stdlib.OrFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/bytes.go
	"byteslen":        stdlib.BytesLenFunc,
	"bytesslice":      stdlib.BytesSliceFunc,

        // https://github.com/zclconf/go-cty/cty/function/stdlib/format.go
	"format":          stdlib.FormatFunc,
	"formatlist":      stdlib.FormatListFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/general.go
	"equal":           stdlib.EqualFunc,
	"notequal":        stdlib.NotEqualFunc,
	"coalesce":        stdlib.CoalesceFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/collection.go
	"hasindex":        stdlib.HasIndexFunc,
	"index":           stdlib.IndexFunc,
	"length":          stdlib.LengthFunc,
	"element":         stdlib.ElementFunc,
	"coalescelist":    stdlib.CoalesceListFunc,
	"compact":         stdlib.CompactFunc,
	"contains":        stdlib.ContainsFunc,
	"distinct":        stdlib.DistinctFunc,
	"chunklist":       stdlib.ChunklistFunc,
	"flatten":         stdlib.FlattenFunc,
	"keys":            stdlib.KeysFunc,
	"lookup":          stdlib.LookupFunc,
	"merge":           stdlib.MergeFunc,
	"reverselist":     stdlib.ReverseListFunc,
	"setproduct":      stdlib.SetProductFunc,
	"slice":           stdlib.SliceFunc,
	"values":          stdlib.ValuesFunc,
	"zipmap":          stdlib.ZipmapFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/csv.go
	"csvdecode":       stdlib.CSVDecodeFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/datetime.go
	"formatdate":      stdlib.FormatDateFunc,
	"timeadd":         stdlib.TimeAddFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/json.go
	"jsondecode":      stdlib.JSONDecodeFunc,
	"jsonencode":      stdlib.JSONEncodeFunc,

	// https://github.com/zclconf/go-cty/cty/function/stdlib/number.go
	"abs":             stdlib.AbsoluteFunc,
	"add":             stdlib.AddFunc,
	"subtract":        stdlib.SubtractFunc,
	"multiply":        stdlib.MultiplyFunc,
	"divide":          stdlib.DivideFunc,
	"modulo":          stdlib.ModuloFunc,
	"greaterthan":     stdlib.GreaterThanFunc,
	"greaterthanorequalto": stdlib.GreaterThanOrEqualToFunc,
	"lessthan":        stdlib.LessThanFunc,
	"lessthanorequalto": stdlib.LessThanOrEqualToFunc,
	"negate":          stdlib.NegateFunc,
	"min":             stdlib.MinFunc,
	"max":             stdlib.MaxFunc,
	"int":             stdlib.IntFunc,
	"ceil":            stdlib.CeilFunc,
	"floor":           stdlib.FloorFunc,
	"log":             stdlib.LogFunc,
	"pow":             stdlib.PowFunc,
	"signum":          stdlib.SignumFunc,
	"parseint":        stdlib.ParseIntFunc,
	// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/regexp.go
	"regex":           stdlib.RegexFunc,
	"regexall":        stdlib.RegexAllFunc,
	// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/sequence.go
	"concat":          stdlib.ConcatFunc,
	"range":           stdlib.RangeFunc,
	// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/set.go
 	//"sethaselement":   stdlib.SetHasElement,
 	//"setunion":        stdlib.SetUnion,
 	//"setintersection": stdlib.SetIntersection,
	//"setsubtract":     stdlib.SetSubtract,
	//"setsymmetricdifference": stdlib.SetSymmetricDifference,
	// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/string.go
	"upper":           stdlib.UpperFunc,
	"lower":           stdlib.LowerFunc,
	"reverse":         stdlib.ReverseFunc,
	"strlen":          stdlib.StrlenFunc,
	"substr":          stdlib.SubstrFunc,
	"join":            stdlib.JoinFunc,
 	"sort":            stdlib.SortFunc,
 	"split":           stdlib.SplitFunc,
 	"chomp":           stdlib.ChompFunc,
 	"indent":          stdlib.IndentFunc,
 	"title":           stdlib.TitleFunc,
 	"trimspace":       stdlib.TrimSpaceFunc,
 	"trim":            stdlib.TrimFunc,
 	"trimprefix":      stdlib.TrimPrefixFunc,
	"trimsuffix":      stdlib.TrimSuffixFunc,
	// https://github.com/zclconf/go-cty/blob/main/cty/function/stdlib/string_replace.go
	"replace":         stdlib.ReplaceFunc,
	"regexreplace":    stdlib.RegexReplaceFunc,
}
