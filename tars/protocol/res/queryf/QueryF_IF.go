//Package queryf comment
// This file war generated by tars2go 1.1
// Generated from QueryF.tars
package queryf

import (
	"context"
	"fmt"

	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/endpointf"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//QueryF struct
type QueryF struct {
	s m.Servant
}

//FindObjectById is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectById(Id string, _opt ...map[string]string) (ret []endpointf.EndpointF, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err, _, ty = _is.SkipToNoCheck(0, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		ret = make([]endpointf.EndpointF, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = ret[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdWithContext(ctx context.Context, Id string, _opt ...map[string]string) (ret []endpointf.EndpointF, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err, _, ty = _is.SkipToNoCheck(0, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		ret = make([]endpointf.EndpointF, length, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = ret[i1].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectById4Any is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectById4Any(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById4Any", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i2, e2 := int32(0), length; i2 < e2; i2++ {

			err = (*ActiveEp)[i2].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i3, e3 := int32(0), length; i3 < e3; i3++ {

			err = (*InactiveEp)[i3].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectById4AnyWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectById4AnyWithContext(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById4Any", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i4, e4 := int32(0), length; i4 < e4; i4++ {

			err = (*ActiveEp)[i4].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i5, e5 := int32(0), length; i5 < e5; i5++ {

			err = (*InactiveEp)[i5].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectById4All is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectById4All(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById4All", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i6, e6 := int32(0), length; i6 < e6; i6++ {

			err = (*ActiveEp)[i6].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i7, e7 := int32(0), length; i7 < e7; i7++ {

			err = (*InactiveEp)[i7].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectById4AllWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectById4AllWithContext(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectById4All", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i8, e8 := int32(0), length; i8 < e8; i8++ {

			err = (*ActiveEp)[i8].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i9, e9 := int32(0), length; i9 < e9; i9++ {

			err = (*InactiveEp)[i9].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameGroup is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameGroup(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameGroup", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i10, e10 := int32(0), length; i10 < e10; i10++ {

			err = (*ActiveEp)[i10].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i11, e11 := int32(0), length; i11 < e11; i11++ {

			err = (*InactiveEp)[i11].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameGroupWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameGroupWithContext(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameGroup", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i12, e12 := int32(0), length; i12 < e12; i12++ {

			err = (*ActiveEp)[i12].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i13, e13 := int32(0), length; i13 < e13; i13++ {

			err = (*InactiveEp)[i13].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameStation is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameStation(Id string, SStation string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(SStation, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameStation", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i14, e14 := int32(0), length; i14 < e14; i14++ {

			err = (*ActiveEp)[i14].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i15, e15 := int32(0), length; i15 < e15; i15++ {

			err = (*InactiveEp)[i15].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameStationWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameStationWithContext(ctx context.Context, Id string, SStation string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(SStation, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameStation", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i16, e16 := int32(0), length; i16 < e16; i16++ {

			err = (*ActiveEp)[i16].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i17, e17 := int32(0), length; i17 < e17; i17++ {

			err = (*InactiveEp)[i17].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameSet is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameSet(Id string, SetId string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(SetId, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameSet", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i18, e18 := int32(0), length; i18 < e18; i18++ {

			err = (*ActiveEp)[i18].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i19, e19 := int32(0), length; i19 < e19; i19++ {

			err = (*InactiveEp)[i19].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//FindObjectByIdInSameSetWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *QueryF) FindObjectByIdInSameSetWithContext(ctx context.Context, Id string, SetId string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Id, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(SetId, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "findObjectByIdInSameSet", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*ActiveEp) = make([]endpointf.EndpointF, length, length)
		for i20, e20 := int32(0), length; i20 < e20; i20++ {

			err = (*ActiveEp)[i20].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	err, _, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return ret, err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return ret, err
		}
		(*InactiveEp) = make([]endpointf.EndpointF, length, length)
		for i21, e21 := int32(0), length; i21 < e21; i21++ {

			err = (*InactiveEp)[i21].ReadBlock(_is, 0, false)
			if err != nil {
				return ret, err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return ret, err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return ret, err
		}
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *QueryF) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *QueryF) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

type _impQueryF interface {
	FindObjectById(Id string) (ret []endpointf.EndpointF, err error)
	FindObjectById4Any(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectById4All(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameGroup(Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameStation(Id string, SStation string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameSet(Id string, SetId string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
}
type _impQueryFWithContext interface {
	FindObjectById(ctx context.Context, Id string) (ret []endpointf.EndpointF, err error)
	FindObjectById4Any(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectById4All(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameGroup(ctx context.Context, Id string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameStation(ctx context.Context, Id string, SStation string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
	FindObjectByIdInSameSet(ctx context.Context, Id string, SetId string, ActiveEp *[]endpointf.EndpointF, InactiveEp *[]endpointf.EndpointF) (ret int32, err error)
}

func findObjectById(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectById(Id)
		if err != nil {
			return err
		}

		err = _os.WriteHead(codec.LIST, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(ret)), 0)
		if err != nil {
			return err
		}
		for _, v := range ret {

			err = v.WriteBlock(_os, 0)
			if err != nil {
				return err
			}
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectById(ctx, Id)
		if err != nil {
			return err
		}

		err = _os.WriteHead(codec.LIST, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(ret)), 0)
		if err != nil {
			return err
		}
		for _, v := range ret {

			err = v.WriteBlock(_os, 0)
			if err != nil {
				return err
			}
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func findObjectById4Any(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	var ActiveEp []endpointf.EndpointF
	var InactiveEp []endpointf.EndpointF
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectById4Any(Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectById4Any(ctx, Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(ActiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range ActiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(InactiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range InactiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func findObjectById4All(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	var ActiveEp []endpointf.EndpointF
	var InactiveEp []endpointf.EndpointF
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectById4All(Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectById4All(ctx, Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(ActiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range ActiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(InactiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range InactiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func findObjectByIdInSameGroup(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	var ActiveEp []endpointf.EndpointF
	var InactiveEp []endpointf.EndpointF
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectByIdInSameGroup(Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectByIdInSameGroup(ctx, Id, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(ActiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range ActiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(InactiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range InactiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func findObjectByIdInSameStation(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	var SStation string
	err = _is.Read_string(&SStation, 2, true)
	if err != nil {
		return err
	}
	var ActiveEp []endpointf.EndpointF
	var InactiveEp []endpointf.EndpointF
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectByIdInSameStation(Id, SStation, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectByIdInSameStation(ctx, Id, SStation, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(ActiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range ActiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(InactiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range InactiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func findObjectByIdInSameSet(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Id string
	err = _is.Read_string(&Id, 1, true)
	if err != nil {
		return err
	}
	var SetId string
	err = _is.Read_string(&SetId, 2, true)
	if err != nil {
		return err
	}
	var ActiveEp []endpointf.EndpointF
	var InactiveEp []endpointf.EndpointF
	if withContext == false {
		_imp := _val.(_impQueryF)
		ret, err := _imp.FindObjectByIdInSameSet(Id, SetId, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impQueryFWithContext)
		ret, err := _imp.FindObjectByIdInSameSet(ctx, Id, SetId, &ActiveEp, &InactiveEp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(ActiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range ActiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(InactiveEp)), 0)
	if err != nil {
		return err
	}
	for _, v := range InactiveEp {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *QueryF) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "findObjectById":
		err := findObjectById(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "findObjectById4Any":
		err := findObjectById4Any(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "findObjectById4All":
		err := findObjectById4All(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "findObjectByIdInSameGroup":
		err := findObjectByIdInSameGroup(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "findObjectByIdInSameStation":
		err := findObjectByIdInSameStation(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "findObjectByIdInSameSet":
		err := findObjectByIdInSameSet(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
