package tf

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"

)

//THIS CODE WAS GENERATED WITH A PYTHON NOTEBOOK

func coerce2RankTypeTensorFlow(inp interface{}, ranks []int64, typ string) (*tf.Tensor, error) {
switch r := ranks; {
	case len(r) == 1 && typ == "DT_DOUBLE":
		output, err := tf.NewTensor(coerceToRank1Float64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_DOUBLE":
		output, err := tf.NewTensor(coerceToRank2Float64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_DOUBLE":
		output, err := tf.NewTensor(coerceToRank3Float64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_DOUBLE":
		output, err := tf.NewTensor(coerceToRank4Float64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_DOUBLE":
		output, err := tf.NewTensor(coerceToRank5Float64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_FLOAT":
		output, err := tf.NewTensor(coerceToRank1Float32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_FLOAT":
		output, err := tf.NewTensor(coerceToRank2Float32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_FLOAT":
		output, err := tf.NewTensor(coerceToRank3Float32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_FLOAT":
		output, err := tf.NewTensor(coerceToRank4Float32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_FLOAT":
		output, err := tf.NewTensor(coerceToRank5Float32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_INT8":
		output, err := tf.NewTensor(coerceToRank1Int8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_INT8":
		output, err := tf.NewTensor(coerceToRank2Int8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_INT8":
		output, err := tf.NewTensor(coerceToRank3Int8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_INT8":
		output, err := tf.NewTensor(coerceToRank4Int8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_INT8":
		output, err := tf.NewTensor(coerceToRank5Int8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_INT16":
		output, err := tf.NewTensor(coerceToRank1Int16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_INT16":
		output, err := tf.NewTensor(coerceToRank2Int16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_INT16":
		output, err := tf.NewTensor(coerceToRank3Int16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_INT16":
		output, err := tf.NewTensor(coerceToRank4Int16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_INT16":
		output, err := tf.NewTensor(coerceToRank5Int16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_INT32":
		output, err := tf.NewTensor(coerceToRank1Int32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_INT32":
		output, err := tf.NewTensor(coerceToRank2Int32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_INT32":
		output, err := tf.NewTensor(coerceToRank3Int32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_INT32":
		output, err := tf.NewTensor(coerceToRank4Int32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_INT32":
		output, err := tf.NewTensor(coerceToRank5Int32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_INT64":
		output, err := tf.NewTensor(coerceToRank1Int64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_INT64":
		output, err := tf.NewTensor(coerceToRank2Int64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_INT64":
		output, err := tf.NewTensor(coerceToRank3Int64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_INT64":
		output, err := tf.NewTensor(coerceToRank4Int64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_INT64":
		output, err := tf.NewTensor(coerceToRank5Int64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_STRING":
		output, err := tf.NewTensor(coerceToRank1String(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_STRING":
		output, err := tf.NewTensor(coerceToRank2String(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_STRING":
		output, err := tf.NewTensor(coerceToRank3String(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_STRING":
		output, err := tf.NewTensor(coerceToRank4String(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_STRING":
		output, err := tf.NewTensor(coerceToRank5String(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_BOOL":
		output, err := tf.NewTensor(coerceToRank1Bool(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_BOOL":
		output, err := tf.NewTensor(coerceToRank2Bool(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_BOOL":
		output, err := tf.NewTensor(coerceToRank3Bool(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_BOOL":
		output, err := tf.NewTensor(coerceToRank4Bool(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_BOOL":
		output, err := tf.NewTensor(coerceToRank5Bool(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_UINT8":
		output, err := tf.NewTensor(coerceToRank1Uint8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_UINT8":
		output, err := tf.NewTensor(coerceToRank2Uint8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_UINT8":
		output, err := tf.NewTensor(coerceToRank3Uint8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_UINT8":
		output, err := tf.NewTensor(coerceToRank4Uint8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_UINT8":
		output, err := tf.NewTensor(coerceToRank5Uint8(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_UINT16":
		output, err := tf.NewTensor(coerceToRank1Uint16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_UINT16":
		output, err := tf.NewTensor(coerceToRank2Uint16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_UINT16":
		output, err := tf.NewTensor(coerceToRank3Uint16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_UINT16":
		output, err := tf.NewTensor(coerceToRank4Uint16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_UINT16":
		output, err := tf.NewTensor(coerceToRank5Uint16(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_UINT32":
		output, err := tf.NewTensor(coerceToRank1Uint32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_UINT32":
		output, err := tf.NewTensor(coerceToRank2Uint32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_UINT32":
		output, err := tf.NewTensor(coerceToRank3Uint32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_UINT32":
		output, err := tf.NewTensor(coerceToRank4Uint32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_UINT32":
		output, err := tf.NewTensor(coerceToRank5Uint32(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 1 && typ == "DT_UINT64":
		output, err := tf.NewTensor(coerceToRank1Uint64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 2 && typ == "DT_UINT64":
		output, err := tf.NewTensor(coerceToRank2Uint64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 3 && typ == "DT_UINT64":
		output, err := tf.NewTensor(coerceToRank3Uint64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 4 && typ == "DT_UINT64":
		output, err := tf.NewTensor(coerceToRank4Uint64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	case len(r) == 5 && typ == "DT_UINT64":
		output, err := tf.NewTensor(coerceToRank5Uint64(inp))
		if err != nil {
			return nil, fmt.Errorf("problem converting feature to tensorflow object:%s", err)
		}
		return output, nil

	
	default:
		data, err := checkDataTypes(inp, r, typ, "a feature")
		if err != nil {
			return nil, err
		}
		output, err := tf.NewTensor(data)
		if err != nil {
			return nil, err
		}
		return output, nil
	}
	return nil, fmt.Errorf("no type or rank in conversion function")
}
func coerceToRank1Float64(inp interface{}) []float64 {
	var output []float64
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(float64))
	}
	return output
}

func coerceToRank2Float64(inp interface{}) [][]float64 {
	var output [][]float64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Float64(val))
	}
	return output
}

func coerceToRank3Float64(inp interface{}) [][][]float64 {
	var output [][][]float64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Float64(val))
	}
	return output
}

func coerceToRank4Float64(inp interface{}) [][][][]float64 {
	var output [][][][]float64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Float64(val))
	}
	return output
}

func coerceToRank5Float64(inp interface{}) [][][][][]float64 {
	var output [][][][][]float64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Float64(val))
	}
	return output
}

func coerceToRank1Float32(inp interface{}) []float32 {
	var output []float32
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(float32))
	}
	return output
}

func coerceToRank2Float32(inp interface{}) [][]float32 {
	var output [][]float32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Float32(val))
	}
	return output
}

func coerceToRank3Float32(inp interface{}) [][][]float32 {
	var output [][][]float32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Float32(val))
	}
	return output
}

func coerceToRank4Float32(inp interface{}) [][][][]float32 {
	var output [][][][]float32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Float32(val))
	}
	return output
}

func coerceToRank5Float32(inp interface{}) [][][][][]float32 {
	var output [][][][][]float32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Float32(val))
	}
	return output
}

func coerceToRank1Int8(inp interface{}) []int8 {
	var output []int8
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(int8))
	}
	return output
}

func coerceToRank2Int8(inp interface{}) [][]int8 {
	var output [][]int8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Int8(val))
	}
	return output
}

func coerceToRank3Int8(inp interface{}) [][][]int8 {
	var output [][][]int8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Int8(val))
	}
	return output
}

func coerceToRank4Int8(inp interface{}) [][][][]int8 {
	var output [][][][]int8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Int8(val))
	}
	return output
}

func coerceToRank5Int8(inp interface{}) [][][][][]int8 {
	var output [][][][][]int8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Int8(val))
	}
	return output
}

func coerceToRank1Int16(inp interface{}) []int16 {
	var output []int16
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(int16))
	}
	return output
}

func coerceToRank2Int16(inp interface{}) [][]int16 {
	var output [][]int16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Int16(val))
	}
	return output
}

func coerceToRank3Int16(inp interface{}) [][][]int16 {
	var output [][][]int16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Int16(val))
	}
	return output
}

func coerceToRank4Int16(inp interface{}) [][][][]int16 {
	var output [][][][]int16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Int16(val))
	}
	return output
}

func coerceToRank5Int16(inp interface{}) [][][][][]int16 {
	var output [][][][][]int16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Int16(val))
	}
	return output
}

func coerceToRank1Int32(inp interface{}) []int32 {
	var output []int32
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(int32))
	}
	return output
}

func coerceToRank2Int32(inp interface{}) [][]int32 {
	var output [][]int32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Int32(val))
	}
	return output
}

func coerceToRank3Int32(inp interface{}) [][][]int32 {
	var output [][][]int32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Int32(val))
	}
	return output
}

func coerceToRank4Int32(inp interface{}) [][][][]int32 {
	var output [][][][]int32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Int32(val))
	}
	return output
}

func coerceToRank5Int32(inp interface{}) [][][][][]int32 {
	var output [][][][][]int32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Int32(val))
	}
	return output
}

func coerceToRank1Int64(inp interface{}) []int64 {
	var output []int64
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(int64))
	}
	return output
}

func coerceToRank2Int64(inp interface{}) [][]int64 {
	var output [][]int64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Int64(val))
	}
	return output
}

func coerceToRank3Int64(inp interface{}) [][][]int64 {
	var output [][][]int64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Int64(val))
	}
	return output
}

func coerceToRank4Int64(inp interface{}) [][][][]int64 {
	var output [][][][]int64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Int64(val))
	}
	return output
}

func coerceToRank5Int64(inp interface{}) [][][][][]int64 {
	var output [][][][][]int64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Int64(val))
	}
	return output
}

func coerceToRank1String(inp interface{}) []string {
	var output []string
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(string))
	}
	return output
}

func coerceToRank2String(inp interface{}) [][]string {
	var output [][]string
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1String(val))
	}
	return output
}

func coerceToRank3String(inp interface{}) [][][]string {
	var output [][][]string
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2String(val))
	}
	return output
}

func coerceToRank4String(inp interface{}) [][][][]string {
	var output [][][][]string
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3String(val))
	}
	return output
}

func coerceToRank5String(inp interface{}) [][][][][]string {
	var output [][][][][]string
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4String(val))
	}
	return output
}

func coerceToRank1Bool(inp interface{}) []bool {
	var output []bool
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(bool))
	}
	return output
}

func coerceToRank2Bool(inp interface{}) [][]bool {
	var output [][]bool
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Bool(val))
	}
	return output
}

func coerceToRank3Bool(inp interface{}) [][][]bool {
	var output [][][]bool
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Bool(val))
	}
	return output
}

func coerceToRank4Bool(inp interface{}) [][][][]bool {
	var output [][][][]bool
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Bool(val))
	}
	return output
}

func coerceToRank5Bool(inp interface{}) [][][][][]bool {
	var output [][][][][]bool
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Bool(val))
	}
	return output
}

func coerceToRank1Uint8(inp interface{}) []uint8 {
	var output []uint8
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(uint8))
	}
	return output
}

func coerceToRank2Uint8(inp interface{}) [][]uint8 {
	var output [][]uint8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Uint8(val))
	}
	return output
}

func coerceToRank3Uint8(inp interface{}) [][][]uint8 {
	var output [][][]uint8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Uint8(val))
	}
	return output
}

func coerceToRank4Uint8(inp interface{}) [][][][]uint8 {
	var output [][][][]uint8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Uint8(val))
	}
	return output
}

func coerceToRank5Uint8(inp interface{}) [][][][][]uint8 {
	var output [][][][][]uint8
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Uint8(val))
	}
	return output
}

func coerceToRank1Uint16(inp interface{}) []uint16 {
	var output []uint16
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(uint16))
	}
	return output
}

func coerceToRank2Uint16(inp interface{}) [][]uint16 {
	var output [][]uint16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Uint16(val))
	}
	return output
}

func coerceToRank3Uint16(inp interface{}) [][][]uint16 {
	var output [][][]uint16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Uint16(val))
	}
	return output
}

func coerceToRank4Uint16(inp interface{}) [][][][]uint16 {
	var output [][][][]uint16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Uint16(val))
	}
	return output
}

func coerceToRank5Uint16(inp interface{}) [][][][][]uint16 {
	var output [][][][][]uint16
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Uint16(val))
	}
	return output
}

func coerceToRank1Uint32(inp interface{}) []uint32 {
	var output []uint32
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(uint32))
	}
	return output
}

func coerceToRank2Uint32(inp interface{}) [][]uint32 {
	var output [][]uint32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Uint32(val))
	}
	return output
}

func coerceToRank3Uint32(inp interface{}) [][][]uint32 {
	var output [][][]uint32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Uint32(val))
	}
	return output
}

func coerceToRank4Uint32(inp interface{}) [][][][]uint32 {
	var output [][][][]uint32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Uint32(val))
	}
	return output
}

func coerceToRank5Uint32(inp interface{}) [][][][][]uint32 {
	var output [][][][][]uint32
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Uint32(val))
	}
	return output
}

func coerceToRank1Uint64(inp interface{}) []uint64 {
	var output []uint64
	for _, val := range inp.([]interface{}) {
		output = append(output, val.(uint64))
	}
	return output
}

func coerceToRank2Uint64(inp interface{}) [][]uint64 {
	var output [][]uint64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank1Uint64(val))
	}
	return output
}

func coerceToRank3Uint64(inp interface{}) [][][]uint64 {
	var output [][][]uint64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank2Uint64(val))
	}
	return output
}

func coerceToRank4Uint64(inp interface{}) [][][][]uint64 {
	var output [][][][]uint64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank3Uint64(val))
	}
	return output
}

func coerceToRank5Uint64(inp interface{}) [][][][][]uint64 {
	var output [][][][][]uint64
	for _, val := range inp.([]interface{}) {
		output = append(output, coerceToRank4Uint64(val))
	}
	return output
}

