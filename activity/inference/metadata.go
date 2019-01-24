package inference

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Model      string        `md:"model"`
	Features   []interface{} `md:"features"`
	Framework  string        `md:"framework"`
	SigDefName string        `md:"sigDefName"`
	Tag        string        `md:"tag"`
}

func (o *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"model":      o.Model,
		"features":   o.Features,
		"framework":  o.Framework,
		"sigDefName": o.SigDefName,
		"tag":        o.Tag,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {

	var err error
	o.Model, err = coerce.ToString(values["model"])
	if err != nil {
		return err
	}
	o.Features, err = coerce.ToArray(values["features"])
	if err != nil {
		return err
	}
	o.Framework, err = coerce.ToString(values["framework"])
	if err != nil {
		return err
	}
	o.SigDefName, err = coerce.ToString(values["sigDefName"])
	if err != nil {
		return err
	}
	o.Tag, err = coerce.ToString(values["tag"])
	if err != nil {
		return err
	}

	return nil
}

type Output struct {
	Result interface{} `md:"result"`
}

func (r *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": r.Result,
	}
}

func (r *Output) FromMap(values map[string]interface{}) error {

	//var err error

	r.Result, _ = values["result"]

	return nil
}
