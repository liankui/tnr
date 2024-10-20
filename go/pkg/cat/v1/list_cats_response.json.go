// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.

package cat

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

func init() {
	core.RegisterJSONTypeDecoder("cat.ListCatsResponse", &ListCatsResponseCodec{})
	core.RegisterJSONTypeEncoder("cat.ListCatsResponse", &ListCatsResponseCodec{})
}

type ListCatsResponseCodec struct {
}

func (codec *ListCatsResponseCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	listCatsResponse := (*ListCatsResponse)(ptr)
	if any.ValueType() == jsoniter.ArrayValue {
		any.ToVal(&listCatsResponse.Cats)
	}
}

func (codec *ListCatsResponseCodec) IsEmpty(ptr unsafe.Pointer) bool {
	listCatsResponse := (*ListCatsResponse)(ptr)
	return listCatsResponse == nil || len(listCatsResponse.Cats) == 0
}

func (codec *ListCatsResponseCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	listCatsResponse := (*ListCatsResponse)(ptr)
	if len(listCatsResponse.Cats) > 0 {
		stream.WriteVal(listCatsResponse.Cats)
	}
}