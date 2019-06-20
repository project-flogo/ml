package inference

import (
	"fmt"
	"strings"
	"sync"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/ml/activity/inference/framework"
	"github.com/project-flogo/ml/activity/inference/framework/tf"
	"github.com/project-flogo/ml/activity/inference/model"
)

var _ tf.TensorflowModel

// variables needed to persist model between inferences
var tfmodelmap map[string]*model.Model
var modelRunMutex sync.Mutex

const (
	ivModel     = "model"
	ivInputName = "inputName"
	ivFeatures  = "features"
	ivFramework = "framework"

	ivSigDef = "sigDefName"
	ivTag    = "tag"

	ovResult = "result"
)

var activityMd = activity.ToMetadata(&Input{}, &Output{})

func init() {
	activity.Register(&Activity{}, New)
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	act := &Activity{}
	logger := ctx.Logger()
	logger.Info("Initialize new activity")
	return act, nil
}

type Settings struct {
}

// InferenceActivity is an Activity that is used to invoke a a ML Model using flogo-ml framework
type Activity struct {
}

// NewActivity creates a new InferenceActivity

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *Activity) Eval(context activity.Context) (done bool, err error) {
	log := context.Logger()
	modelName := context.GetInput(ivModel).(string)
	features := context.GetInput(ivFeatures).([]interface{})
	fw := context.GetInput(ivFramework).(string)

	tfFramework := framework.Get(fw)
	log.Debugf("Using Framework %s", tfFramework.FrameworkTyp())
	if tfFramework == nil {
		log.Errorf("%s framework not registered", fw)

		return true, fmt.Errorf("%s framework not registered", fw)
	}
	log.Debugf("Loaded Framework: " + tfFramework.FrameworkTyp())

	// Defining the flags to be used to load model
	flags := model.ModelFlags{
		Tag:    context.GetInput("tag").(string),
		SigDef: context.GetInput("sigDefName").(string),
	}

	// if modelmap does not exist then make it
	if tfmodelmap == nil {
		tfmodelmap = make(map[string]*model.Model)
		log.Info("Making map of models with keys of 'ModelKey'.")
	}

	// check if this instance of tf model already exists if not load it
	modelKey := context.ActivityHost().Name() + ":" + context.Name() + context.GetInput("model").(string)
	log.Info("ModelKey is:", modelKey)
	if tfmodelmap[modelKey] == nil {
		tfmodelmap[modelKey], err = model.Load(modelName, tfFramework, flags)
		if err != nil {
			return true, err
		}

	} else {
		log.Debug("Model already loaded - skipping loading")
	}

	log.Debug("Incoming features: ")
	log.Debug(features)

	// Grab the input feature set and parse out the feature labels and values
	inputSample := make(map[string]interface{})
	for _, feat := range features {
		featMap := feat.(map[string]interface{})
		inputName := featMap["name"].(string)
		inmodel := false
		for key := range tfmodelmap[modelKey].Metadata.Inputs.Features{
			if key==inputName{
				inputSample[inputName] = featMap["data"]
				inmodel=true
			}
		}
		if !inmodel{
			return false,fmt.Errorf("%s not an input into model",featMap["name"].(string))
		}
	}

	log.Info("Parsing of features completed")

	modelRunMutex.Lock()
	tfmodelmap[modelKey].SetInputs(inputSample)
	output, err := tfmodelmap[modelKey].Run(tfFramework)
	modelRunMutex.Unlock()
	if err != nil {
		log.Errorf("Error running the ml framework: %s", err)
		return true, err
	}

	log.Debug("Model execution completed with result:")
	log.Debug(output)


	if strings.Contains(tfmodelmap[modelKey].Metadata.Method, "tensorflow/serving/classify") {
		var out = make(map[string]interface{})

		classes := output["classes"].([][]string)[0]
		scores := output["scores"].([][]float32)[0]

		for i := 0; i < len(classes); i++ {
			out[classes[i]] = scores[i]
		}

		context.SetOutput(ovResult, out)
	} else {
		context.SetOutput(ovResult, output)
	}

	return true, nil
}
