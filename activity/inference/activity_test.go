package inference

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support/test"
	"github.com/project-flogo/ml/activity/inference/framework/tf"
	"github.com/stretchr/testify/assert"
	tfactual "github.com/tensorflow/tensorflow/tensorflow/go"
)

var _ tf.TensorflowModel

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestDNNEstimator(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test of Estimator DNN Regressor model
	fmt.Println("Unit test of Estimator Regressor model")
	tc.SetInput("model", "testModels/Archive_estDNNrgr.zip")
	tc.SetInput("inputName", "inputs")
	var estInputsB = make(map[string]interface{})
	estInputsB["one"] = 0.140586
	estInputsB["two"] = 0.140586
	estInputsB["three"] = 0.140586
	estInputsB["label"] = 0.

	var featuresB []interface{}
	featuresB = append(featuresB, map[string]interface{}{
		"name": "inputs",
		"data": estInputsB,
	})

	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", featuresB)
	fmt.Println("blah")

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["outputs"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}
	fmt.Println(tc.GetOutput("result"))
}

func TestEstimatorLinearRegressor(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test of Estimator Linear Regressor model
	fmt.Println("Unit test of Linear Regressor Estimator model")
	tc.SetInput("model", "testModels/Archive_LinReg.zip")
	tc.SetInput("inputName", "inputs")
	var estInputsC = make(map[string]interface{})
	estInputsC["one"] = 0.140586
	estInputsC["two"] = 0.140586
	estInputsC["three"] = 0.140586
	estInputsC["label"] = 0.

	var featuresC []interface{}
	featuresC = append(featuresC, map[string]interface{}{
		"name": "inputs",
		"data": estInputsC,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", featuresC)

	done, err = act.Eval(tc)
	if done == false {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}
	fmt.Println(tc.GetOutput("result"))
}

func TestPairwaiseMul(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test of Pairwaise Multiplication model
	fmt.Println("Unit test of Pairwaise Multiplication model")
	tc.SetInput("model", "testModels/Archive_pairwise_multi.zip")

	var features2 []interface{}
	features2 = append(features2, map[string]interface{}{
		"name": "X1",
		"data": [][]float32{{0.23, 4.5, -3.1}, {7.1, 3.14159, -0.00123}},
	})
	features2 = append(features2, map[string]interface{}{
		"name": "X2",
		"data": [][]float32{{4.34782608, 0.2222222222, -0.3225806451612903},
			{0.14084507042253522, 0.31831015504887655, -813.0081300813008}},
	})

	fmt.Println(features2)

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features2)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["pred"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}
	fmt.Println(tc.GetOutput("result"))
}

func TestCNNModel(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit test of simple CNN model")
	tc.SetInput("model", "testModels/Archive_simpleCNN.zip")

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "X",
		"data": [][][][]float32{
			{{{0.0000000856947568}}, {{0.00000331318370}}, {{0.0000858655563}}, {{0.00149167657}}, {{0.0173705094}}, {{0.135591557}}, {{0.709471493}}, {{2.48839579}}, {{5.85040827}}, {{9.22008867}}},
			{{{9.22008867}}, {{5.85040827}}, {{2.48839579}}, {{00.709471493}}, {{0.135591557}}, {{0.00149167657}}, {{0.0000858655563}}, {{0.00000331318370}}, {{0.0000000856947568}}, {{0.}}},
			{{{0.0173705094}}, {{0.135591557}}, {{0.709471493}}, {{2.48839579}}, {{5.85040827}}, {{9.22008867}}, {{5.85040827}}, {{2.48839579}}, {{0.709471493}}, {{0.135591557}}},
		},
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["pred"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

	//check result attr
	fmt.Println(tc.GetOutput("result"))
}

func TestByteModel(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit test of passing byte to model")
	tc.SetInput("model", "testModels/Archive_byte_test_model.zip")

	d := []int32{1, 2, 3, 4, 5, 10}
	data, err := tfactual.NewTensor(d)
	if err != nil {
		fmt.Println("problem loading TF Tensor")
	}

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "inputs",
		"data": data,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["out"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

	//check result attr
	fmt.Println(tc.GetOutput("result"))
}

func TestEstimatorClassifier(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test of Estimator Classifier model
	fmt.Println("Unit test of Estimator Classifier model")
	tc.SetInput("model", "testModels/Archive_estDNNClf.zip")
	tc.SetInput("inputName", "inputs")
	var estInputsA = make(map[string]interface{})
	estInputsA["one"] = 0.140586
	estInputsA["two"] = 0.140586
	estInputsA["three"] = 0.140586
	estInputsA["label"] = 0

	var featuresA []interface{}
	featuresA = append(featuresA, map[string]interface{}{
		"name": "inputs",
		"data": estInputsA,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", featuresA)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["0"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}
	fmt.Println(tc.GetOutput("result"))
}

func TestCreate(t *testing.T) {

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, _ := New(iCtx)

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestCNNModel2(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit test of simple CNN model")
	tc.SetInput("model", "testModels/Archive_simpleCNN.zip")

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "X",
		"data": [][][][]float32{
			{{{0.0000000856947568}}, {{0.00000331318370}}, {{0.0000858655563}}, {{0.00149167657}}, {{0.0173705094}}, {{0.135591557}}, {{0.709471493}}, {{2.48839579}}, {{5.85040827}}, {{9.22008867}}},
			{{{9.22008867}}, {{5.85040827}}, {{2.48839579}}, {{00.709471493}}, {{0.135591557}}, {{0.00149167657}}, {{0.0000858655563}}, {{0.00000331318370}}, {{0.0000000856947568}}, {{0.}}},
			{{{0.0173705094}}, {{0.135591557}}, {{0.709471493}}, {{2.48839579}}, {{5.85040827}}, {{9.22008867}}, {{5.85040827}}, {{2.48839579}}, {{0.709471493}}, {{0.135591557}}},
		},
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["pred"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

	//check result attr
	fmt.Println(tc.GetOutput("result"))
}

func TestEstimatorClassifierTESTING(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test of Estimator Classifier model
	fmt.Println("Unit test of Estimator Classifier model")
	tc.SetInput("model", "testModels/Archive_estDNNClf.zip")
	tc.SetInput("inputName", "inputs")
	var estInputsA = make(map[string]interface{})
	estInputsA["one"] = 0.140586
	estInputsA["two"] = 0.140586
	estInputsA["three"] = 0.140586
	estInputsA["label"] = 0

	var featuresA []interface{}
	featuresA = append(featuresA, map[string]interface{}{
		"name": "inputs",   //"one",
		"data": estInputsA, // 0.140586,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", featuresA)

	testval := make(map[string]interface{})
	testval["0"] = 0.32395765
	testval["1"] = 0.3435985
	testval["2"] = 0.33244383

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["0"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}
}

func TestFloat64Rank5InterfaceIn(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit TestFloat64Rank5 with interface input")
	tc.SetInput("model", "testModels/Archive_float64.zip")

	in := []interface{}{
		[]interface{}{
			[]interface{}{[]interface{}{[]interface{}{0.8931730159926226, 0.5311302512161289}}},
		},
	}

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "inputs",
		"data": in,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["Yout"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

}

func TestFloat64Rank5DoubleIn(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit TestFloat64Rank5 with Double input")
	tc.SetInput("model", "testModels/Archive_float64.zip")

	in := [][][][][]float64{{{{{0.7814656252241947, 0.28796209040538656}}}}}

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "inputs",
		"data": in,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["Yout"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

}

func TestCNNModelInterfaceIn(t *testing.T) {

	var done bool
	var err error

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	iCtx := test.NewActivityInitContext(&Settings{}, mf)
	act, err := New(iCtx)

	tc := test.NewActivityContext(act.Metadata())

	// Unit test ofSimple CNN model
	fmt.Println("Unit test of simple CNN model - with data formated as an interface")
	tc.SetInput("model", "testModels/Archive_simpleCNN.zip")

	// var in interface{}
	in := []interface{}{
		[]interface{}{
			[]interface{}{[]interface{}{0.0000000856947568}},
			[]interface{}{[]interface{}{0.00000331318370}},
			[]interface{}{[]interface{}{0.0000858655563}},
			[]interface{}{[]interface{}{0.00149167657}},
			[]interface{}{[]interface{}{0.0173705094}},
			[]interface{}{[]interface{}{0.135591557}},
			[]interface{}{[]interface{}{0.709471493}},
			[]interface{}{[]interface{}{2.48839579}},
			[]interface{}{[]interface{}{5.85040827}},
			[]interface{}{[]interface{}{9.22008867}},
		},
	}

	var features3 []interface{}
	features3 = append(features3, map[string]interface{}{
		"name": "X",
		"data": in,
	})

	tc.SetInput("inputName", "inputs")
	tc.SetInput("framework", "Tensorflow")
	tc.SetInput("sigDefName", "serving_default")
	tc.SetInput("tag", "serve")
	tc.SetInput("features", features3)

	done, err = act.Eval(tc)
	out := tc.GetOutput("result")
	if out.(map[string]interface{})["pred"] == nil {
		assert.Fail(t, fmt.Sprintf("Error raised: %s", err))
	} else {
		assert.True(t, done, fmt.Sprintf("Evaluation came back: %t", done))
	}

}
