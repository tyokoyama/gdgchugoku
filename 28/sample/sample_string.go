// Code generated by "stringer -type=Sample"; DO NOT EDIT

package main

import "fmt"

const _Sample_name = "hogefugapiyo"

var _Sample_index = [...]uint8{0, 4, 8, 12}

func (i Sample) String() string {
	if i < 0 || i >= Sample(len(_Sample_index)-1) {
		return fmt.Sprintf("Sample(%d)", i)
	}
	return _Sample_name[_Sample_index[i]:_Sample_index[i+1]]
}
