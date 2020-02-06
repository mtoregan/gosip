// Code generated by `ggen -ent RecycleBinItem -helpers Data,Normalized`; DO NOT EDIT.

package api

import (
	"encoding/json"
)

/* Response helpers */

// Data response helper
func (recycleBinItemResp *RecycleBinItemResp) Data() *RecycleBinItemInfo {
	data := NormalizeODataItem(*recycleBinItemResp)
	res := &RecycleBinItemInfo{}
	json.Unmarshal(data, &res)
	return res
}

// Normalized returns normalized body
func (recycleBinItemResp *RecycleBinItemResp) Normalized() []byte {
	return NormalizeODataItem(*recycleBinItemResp)
}
