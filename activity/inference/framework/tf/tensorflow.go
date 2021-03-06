package tf

import (
	"github.com/project-flogo/ml/activity/inference/framework"
)

// TensorflowModel is
type TensorflowModel struct {
	frameworkTyp string
}

func init() {
	instance := new()
	framework.Register(instance)
}

func new() *TensorflowModel {
	return &TensorflowModel{frameworkTyp: "Tensorflow"}
}

func (a *TensorflowModel) FrameworkTyp() string {
	return a.frameworkTyp
}
