package param

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	query paramType = iota
	path
)

type paramType int

type Holder struct {
	ctx       *gin.Context
	paramName string
	paramType paramType
	required  bool
	value     *ValueHodler
}

type ValueHodler struct {
	rawValue     string
	defaultValue string
}

func Query(ctx *gin.Context) *Holder {
	return &Holder{
		ctx:       ctx,
		paramType: query,
		value:     &ValueHodler{},
	}
}

func Path(ctx *gin.Context) *Holder {
	return &Holder{
		ctx:       ctx,
		required:  true,
		paramType: path,
		value:     &ValueHodler{},
	}
}

func (p *Holder) Required() *Holder {
	p.required = true
	return p
}

func (p *Holder) Name(name string) *Holder {
	p.paramName = name
	return p
}

func (p *Holder) Default(value string) *Holder {
	p.value.defaultValue = value
	return p
}

func (p *Holder) Value() *ValueHodler {
	var rawValue string
	var ok bool

	if p.paramName == "" {
		panic("The parameter name must not be empty")
	}

	if p.paramType == query {
		rawValue, ok = p.ctx.GetQuery(p.paramName)
	} else {
		rawValue, ok = p.ctx.Params.Get(p.paramName)
	}

	if p.required && (!ok || rawValue == "") {
		panic("params is absent: " + p.paramName)
	}
	p.value.rawValue = rawValue
	return p.value
}

func (p *ValueHodler) GetInt8() int8 {
	return getTyped[int8](p)
}

func (p *ValueHodler) GetInt16() int16 {
	return getTyped[int16](p)
}

func (p *ValueHodler) GetInt32() int32 {
	return getTyped[int32](p)
}

func (p *ValueHodler) GetInt64() int64 {
	return getTyped[int64](p)
}

func (p *ValueHodler) GetInt() int {
	return getTyped[int](p)
}

func (p *ValueHodler) GetUint() uint {
	return getTyped[uint](p)
}

func (p *ValueHodler) GetUint8() uint8 {
	return getTyped[uint8](p)
}

func (p *ValueHodler) GetUint16() uint16 {
	return getTyped[uint16](p)
}

func (p *ValueHodler) GetUint32() uint32 {
	return getTyped[uint32](p)
}

func (p *ValueHodler) GetUint64() uint64 {
	return getTyped[uint64](p)
}

func (p *ValueHodler) GetString() string {
	return getTyped[string](p)
}

func (p *ValueHodler) GetFloat32() float32 {
	return getTyped[float32](p)
}

func (p *ValueHodler) GetFloat64() float64 {
	return getTyped[float64](p)
}

func (p *ValueHodler) GetBool() bool {
	return getTyped[bool](p)
}

func getTyped[T any](holder *ValueHodler) (res T) {
	var rawVal = holder.rawValue
	var zero T
	var err error

	if rawVal == "" {
		if holder.defaultValue == "" {
			// 返回零值
			return reflect.Zero(reflect.TypeOf(zero)).Interface().(T)
		} else {
			// 使用默认值解析
			rawVal = holder.defaultValue
		}
	}

	// 有值
	var parsedValue any
	switch any(zero).(type) {
	case int:
		parsedValue, err = strconv.Atoi(rawVal)
	case int8:
		parsedValue, err = strconv.ParseInt(rawVal, 10, 8)
	case int16:
		parsedValue, err = strconv.ParseInt(rawVal, 10, 16)
	case int32:
		parsedValue, err = strconv.ParseInt(rawVal, 10, 32)
	case int64:
		parsedValue, err = strconv.ParseInt(rawVal, 10, 64)
	case uint:
		parsedValue, err = strconv.ParseUint(rawVal, 10, 64)
	case uint8:
		parsedValue, err = strconv.ParseUint(rawVal, 10, 8)
	case uint16:
		parsedValue, err = strconv.ParseUint(rawVal, 10, 16)
	case uint32:
		parsedValue, err = strconv.ParseUint(rawVal, 10, 32)
	case uint64:
		parsedValue, err = strconv.ParseUint(rawVal, 10, 64)
	case float32:
		parsedValue, err = strconv.ParseFloat(rawVal, 32)
	case float64:
		parsedValue, err = strconv.ParseFloat(rawVal, 64)
	case string:
		parsedValue = rawVal
	case bool:
		parsedValue, err = strconv.ParseBool(rawVal)
	default:
		err = fmt.Errorf("unknown type %v", reflect.TypeOf(zero))
	}

	if err != nil {
		panic(err)
	}
	return parsedValue.(T)
}
