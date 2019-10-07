package tf

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
)

func buildStructures(data []interface{}, typ string, rank int) (interface{}, error) {
	var out interface{}
	var err error

	if err != nil {
		return nil, fmt.Errorf("unable to convert slice to rank %d %s tensor: %s", rank, typ, err)
	}
	switch typ {
	case "DT_DOUBLE":
		switch rank {
		case 1:
			out, err = iToR1F64(data)
		case 2:
			out, err = iToR2F64(data)
		case 3:
			out, err = iToR3F64(data)
		case 4:
			out, err = iToR4F64(data)
		case 5:
			out, err = iToR5F64(data)
		case 6:
			out, err = iToR6F64(data)
		default:
			out = data
		}
	case "DT_FLOAT":
		switch rank {
		case 1:
			out, err = iToR1F32(data)
		case 2:
			out, err = iToR2F32(data)
		case 3:
			out, err = iToR3F32(data)
		case 4:
			out, err = iToR4F32(data)
		case 5:
			out, err = iToR5F32(data)
		case 6:
			out, err = iToR6F32(data)
		default:
			out = data
		}
	case "DT_STRING":
		switch rank {
		case 1:
			out, err = iToR1STR(data)
		case 2:
			out, err = iToR2STR(data)
		case 3:
			out, err = iToR3STR(data)
		case 4:
			out, err = iToR4STR(data)
		case 5:
			out, err = iToR5STR(data)
		case 6:
			out, err = iToR6STR(data)
		default:
			out = data
		}
	case "DT_INT32", "DT_INT16", "DT_INT8", "DT_UINT8":
		switch rank {
		case 1:
			out, err = iToR1I32(data)
		case 2:
			out, err = iToR2I32(data)
		case 3:
			out, err = iToR3I32(data)
		case 4:
			out, err = iToR4I32(data)
		case 5:
			out, err = iToR5I32(data)
		case 6:
			out, err = iToR6I32(data)
		default:
			out = data
		}
	case "DT_INT64":
		switch rank {
		case 1:
			out, err = iToR1I64(data)
		case 2:
			out, err = iToR2I64(data)
		case 3:
			out, err = iToR3I64(data)
		case 4:
			out, err = iToR4I64(data)
		case 5:
			out, err = iToR5I64(data)
		case 6:
			out, err = iToR6I64(data)
		default:
			out = data
		}
	case "DT_BOOL":
		switch rank {
		case 1:
			out, err = iToR1BOOL(data)
		case 2:
			out, err = iToR2BOOL(data)
		case 3:
			out, err = iToR3BOOL(data)
		case 4:
			out, err = iToR4BOOL(data)
		case 5:
			out, err = iToR5BOOL(data)
		case 6:
			out, err = iToR6BOOL(data)
		default:
			out = data
		}

		//SO FAR IGNORING DT_COMPLEX64,DT_COMPLEX128,DT_RESOURCE,DT_VARIANT,DT_UINT32,DT_UINT64
	default:
		out = data
	}

	if err != nil {
		return nil, fmt.Errorf("unable to convert slice to rank %d %s tensor: %s", rank, typ, err)
	}
	return out, nil
}

func iToR6F64(struc []interface{}) ([][][][][][]float64, error) {
	var out [][][][][][]float64
	var err error

	for _, val := range struc {
		var tmp [][][][][]float64
		tmp, err = iToR5F64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5F64(struc []interface{}) ([][][][][]float64, error) {
	var out [][][][][]float64
	var err error

	for _, val := range struc {
		var tmp [][][][]float64
		tmp, err = iToR4F64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4F64(struc []interface{}) ([][][][]float64, error) {
	var out [][][][]float64
	var err error

	for _, val := range struc {
		var tmp [][][]float64
		tmp, err = iToR3F64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3F64(struc []interface{}) ([][][]float64, error) {
	var out [][][]float64
	var err error

	for _, val := range struc {
		var tmp [][]float64
		tmp, err = iToR2F64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2F64(struc []interface{}) ([][]float64, error) {
	var out [][]float64
	var err error

	for _, val := range struc {
		var tmp []float64
		tmp, err = iToR1F64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1F64(struc []interface{}) ([]float64, error) {
	var out []float64
	var err error
	for _, val := range struc {
		var tmp float64
		tmp, err = iToR0F64(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0F64(struc interface{}) (float64, error) {
	out, err := coerce.ToFloat64(struc)
	return out, err
}

func iToR6F32(struc []interface{}) ([][][][][][]float32, error) {
	var out [][][][][][]float32
	var err error

	for _, val := range struc {
		var tmp [][][][][]float32
		tmp, err = iToR5F32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5F32(struc []interface{}) ([][][][][]float32, error) {
	var out [][][][][]float32
	var err error

	for _, val := range struc {
		var tmp [][][][]float32
		tmp, err = iToR4F32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4F32(struc []interface{}) ([][][][]float32, error) {
	var out [][][][]float32
	var err error

	for _, val := range struc {
		var tmp [][][]float32
		tmp, err = iToR3F32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3F32(struc []interface{}) ([][][]float32, error) {
	var out [][][]float32
	var err error

	for _, val := range struc {
		var tmp [][]float32
		tmp, err = iToR2F32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2F32(struc []interface{}) ([][]float32, error) {
	var out [][]float32
	var err error

	for _, val := range struc {
		var tmp []float32
		tmp, err = iToR1F32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1F32(struc []interface{}) ([]float32, error) {
	var out []float32
	var err error
	for _, val := range struc {
		var tmp float32
		tmp, err = iToR0F32(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0F32(struc interface{}) (float32, error) {
	out, err := coerce.ToFloat32(struc)
	return out, err
}

func iToR6STR(struc []interface{}) ([][][][][][]string, error) {
	var out [][][][][][]string
	var err error

	for _, val := range struc {
		var tmp [][][][][]string
		tmp, err = iToR5STR(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5STR(struc []interface{}) ([][][][][]string, error) {
	var out [][][][][]string
	var err error

	for _, val := range struc {
		var tmp [][][][]string
		tmp, err = iToR4STR(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4STR(struc []interface{}) ([][][][]string, error) {
	var out [][][][]string
	var err error

	for _, val := range struc {
		var tmp [][][]string
		tmp, err = iToR3STR(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3STR(struc []interface{}) ([][][]string, error) {
	var out [][][]string
	var err error

	for _, val := range struc {
		var tmp [][]string
		tmp, err = iToR2STR(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2STR(struc []interface{}) ([][]string, error) {
	var out [][]string
	var err error

	for _, val := range struc {
		var tmp []string
		tmp, err = iToR1STR(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1STR(struc []interface{}) ([]string, error) {
	var out []string
	var err error
	for _, val := range struc {
		var tmp string
		tmp, err = iToR0STR(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0STR(struc interface{}) (string, error) {
	out, err := coerce.ToString(struc)
	return out, err
}

func iToR6I32(struc []interface{}) ([][][][][][]int32, error) {
	var out [][][][][][]int32
	var err error

	for _, val := range struc {
		var tmp [][][][][]int32
		tmp, err = iToR5I32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5I32(struc []interface{}) ([][][][][]int32, error) {
	var out [][][][][]int32
	var err error

	for _, val := range struc {
		var tmp [][][][]int32
		tmp, err = iToR4I32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4I32(struc []interface{}) ([][][][]int32, error) {
	var out [][][][]int32
	var err error

	for _, val := range struc {
		var tmp [][][]int32
		tmp, err = iToR3I32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3I32(struc []interface{}) ([][][]int32, error) {
	var out [][][]int32
	var err error

	for _, val := range struc {
		var tmp [][]int32
		tmp, err = iToR2I32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2I32(struc []interface{}) ([][]int32, error) {
	var out [][]int32
	var err error

	for _, val := range struc {
		var tmp []int32
		tmp, err = iToR1I32(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1I32(struc []interface{}) ([]int32, error) {
	var out []int32
	var err error
	for _, val := range struc {
		var tmp int32
		tmp, err = iToR0I32(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0I32(struc interface{}) (int32, error) {
	out, err := coerce.ToInt32(struc)
	return out, err
}

func iToR6I64(struc []interface{}) ([][][][][][]int64, error) {
	var out [][][][][][]int64
	var err error

	for _, val := range struc {
		var tmp [][][][][]int64
		tmp, err = iToR5I64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5I64(struc []interface{}) ([][][][][]int64, error) {
	var out [][][][][]int64
	var err error

	for _, val := range struc {
		var tmp [][][][]int64
		tmp, err = iToR4I64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4I64(struc []interface{}) ([][][][]int64, error) {
	var out [][][][]int64
	var err error

	for _, val := range struc {
		var tmp [][][]int64
		tmp, err = iToR3I64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3I64(struc []interface{}) ([][][]int64, error) {
	var out [][][]int64
	var err error

	for _, val := range struc {
		var tmp [][]int64
		tmp, err = iToR2I64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2I64(struc []interface{}) ([][]int64, error) {
	var out [][]int64
	var err error

	for _, val := range struc {
		var tmp []int64
		tmp, err = iToR1I64(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1I64(struc []interface{}) ([]int64, error) {
	var out []int64
	var err error
	for _, val := range struc {
		var tmp int64
		tmp, err = iToR0I64(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0I64(struc interface{}) (int64, error) {
	out, err := coerce.ToInt64(struc)
	return out, err
}

func iToR6BOOL(struc []interface{}) ([][][][][][]bool, error) {
	var out [][][][][][]bool
	var err error

	for _, val := range struc {
		var tmp [][][][][]bool
		tmp, err = iToR5BOOL(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR5BOOL(struc []interface{}) ([][][][][]bool, error) {
	var out [][][][][]bool
	var err error

	for _, val := range struc {
		var tmp [][][][]bool
		tmp, err = iToR4BOOL(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR4BOOL(struc []interface{}) ([][][][]bool, error) {
	var out [][][][]bool
	var err error

	for _, val := range struc {
		var tmp [][][]bool
		tmp, err = iToR3BOOL(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR3BOOL(struc []interface{}) ([][][]bool, error) {
	var out [][][]bool
	var err error

	for _, val := range struc {
		var tmp [][]bool
		tmp, err = iToR2BOOL(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR2BOOL(struc []interface{}) ([][]bool, error) {
	var out [][]bool
	var err error

	for _, val := range struc {
		var tmp []bool
		tmp, err = iToR1BOOL(val.([]interface{}))
		out = append(out, tmp)
	}
	return out, err
}

func iToR1BOOL(struc []interface{}) ([]bool, error) {
	var out []bool
	var err error
	for _, val := range struc {
		var tmp bool
		tmp, err = iToR0BOOL(val)
		out = append(out, tmp)
	}

	return out, err
}

func iToR0BOOL(struc interface{}) (bool, error) {
	out, err := coerce.ToBool(struc)
	return out, err
}
