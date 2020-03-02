package tf

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/project-flogo/core/support/log"
	models "github.com/project-flogo/ml/activity/inference/model"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// log is the default package logger

// Run is used to execute a Tensorflow model with the model input data
func (i *TensorflowModel) Run(model *models.Model) (out map[string]interface{}, err error) {
	// Grab native tf SavedModel
	savedModel := model.Instance.(*tf.SavedModel)

	var inputOps = make(map[string]*tf.Operation)
	// var inputOpsType = make(map[string]string)
	var outputOps []tf.Output

	// Validate that the operations exsist and create operation
	for k, v := range model.Metadata.Inputs.Features {
		if validateOperation(v.Name, savedModel) == false {
			return nil, fmt.Errorf("Invalid operation %s", v.Name)
		}
		inputOps[k] = savedModel.Graph.Operation(v.Name)
	}

	// Create output operations
	var outputOrder []string
	for k, o := range model.Metadata.Outputs {
		outputOps = append(outputOps, savedModel.Graph.Operation(o.Name).Output(0))
		outputOrder = append(outputOrder, k)
	}

	// create input tensors and add to map
	inputs := make(map[tf.Output]*tf.Tensor)
	for inputName, inputMap := range inputOps {
		v := reflect.ValueOf(model.Inputs[inputName])
		switch v.Kind() {
		case reflect.Map:
			log.RootLogger().Debug("Data is determined to be a map and is being converted to tf.tensor")
			// Need to check names against pb structure, right now just assume it
			examplePb, err := createInputExampleTensor(model.Inputs[inputName])
			if err != nil {
				return nil, err
			}
			inputs[inputMap.Output(0)] = examplePb

		case reflect.Slice, reflect.Array:
			log.RootLogger().Debug("Data is determined to be a slice/array and is being converted to tf.tensor")

			typ := model.Metadata.Inputs.Features[inputName].Type
			rank := len(model.Metadata.Inputs.Features[inputName].Shape)
			switch d := model.Inputs[inputName].(type) {
			case []interface{}:
				log.RootLogger().Debug("Input is a slice/array (probably something like [][]float64")
				builtstruc, err := buildStructures(model.Inputs[inputName].([]interface{}), typ, rank)
				if err != nil {
					return nil, fmt.Errorf("unable to convert slice to rank %d %s tensor: %s", rank, typ, err)
				}
				inputs[inputMap.Output(0)], err = tf.NewTensor(builtstruc)
				if err != nil {
					return nil, fmt.Errorf(" unable to convert slice to rank %d %s tensor: %s", rank, typ, err)
				}
			default:
				log.RootLogger().Debug("Input is of default type (probably something like [][]float64")
				inputs[inputMap.Output(0)], err = tf.NewTensor(d)
				if err != nil {
					return nil, fmt.Errorf("unable to convert slice to rank %d %s tensor: %s", rank, typ, err)
				}
			}

		case reflect.Ptr:
			log.RootLogger().Debug("Data is determined to be a pointer and is being converted to tf.tensor")
			if val, ok := model.Inputs[inputName].(*tf.Tensor); ok {
				inputs[inputMap.Output(0)] = val
			} else {
				if val2, ok2 := model.Inputs[inputName].(*[]byte); ok2 {
					inputs[inputMap.Output(0)], err = tf.NewTensor(val2)
					if err != nil {
						return nil, err
					}
				} else {
					return nil, fmt.Errorf("Interface not casting to Tensor or byte object. Is your pointer a tensor?")
				}

			}

		default:

			log.RootLogger().Info("Type not a Slice, Array, Map, or Pointer/Tensor, but still trying to make a tf.Tensor.")
			log.RootLogger().Debug("model.Inputs[inputName] = ", model.Inputs[inputName])
			inputs[inputMap.Output(0)], err = tf.NewTensor(model.Inputs[inputName])
			if err != nil {
				return nil, err
			}
		}
	}

	results, err := savedModel.Session.Run(inputs, outputOps, nil)
	if err != nil {
		return nil, err
	}

	// Iterate over the expected outputs, find the actual and map into map
	out = make(map[string]interface{})
	for k := range model.Metadata.Outputs {
		for i := 0; i < len(outputOrder); i++ {
			if outputOrder[i] == k {
				out[k] = getTensorValue(results[i])
			}
		}
	}

	return out, nil

}

func getTensorValue(tensor *tf.Tensor) interface{} {
	switch tensor.Value().(type) {
	case [][]string:
		return tensor.Value().([][]string)
	case []string:
		return tensor.Value().([]string)
	case string:
		return tensor.Value().(string)
	case float32:
		return tensor.Value().(float32)
	case []float32:
		return tensor.Value().([]float32)
	case [][]float32:
		return tensor.Value().([][]float32)
	case []float64:
		return tensor.Value().([]float64)
	case [][]float64:
		return tensor.Value().([][]float64)
	case []int64:
		return tensor.Value().([]int64)
	case [][]int64:
		return tensor.Value().([][]int64)
	case int32:
		return tensor.Value().(int32)
	case []int32:
		return tensor.Value().([]int32)
	case [][]int32:
		return tensor.Value().([][]int32)
	case byte:
		return tensor.Value().(byte)
	case []byte:
		return tensor.Value().([]byte)
	case [][]byte:
		return tensor.Value().([][]byte)
	case []int:
		return tensor.Value().([]int)
	}

	return tensor.Value()
}

func createInputExampleTensor(featMap interface{}) (*tf.Tensor, error) {
	pb, err := Example(featMap.(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("Failed to create Example: %s", err)
	}

	byteList, err := proto.Marshal(pb)
	if err != nil {
		return nil, fmt.Errorf("marshaling error: %s", err)
	}

	newTensor, err := tf.NewTensor([]string{string(byteList)})
	if err != nil {
		return nil, err
	}

	return newTensor, nil
}

func validateOperation(op string, savedModel *tf.SavedModel) bool {

	tfOp := savedModel.Graph.Operation(op)
	if tfOp == nil {
		return false
	}
	return true
}
