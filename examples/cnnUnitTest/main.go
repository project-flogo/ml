package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/project-flogo/contrib/activity/log"
	"github.com/project-flogo/contrib/trigger/rest"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/ml/activity/inference"
)

var (
	httpport = os.Getenv("HTTPPORT")
)

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	// Convert the HTTPPort to an integer
	port, err := strconv.Atoi(httpport)
	if err != nil {
		fmt.Println("Error:", err)
		return app
	}

	// Defining the trigger used - in this case the REST Trigger
	trg := app.NewTrigger(&rest.Trigger{}, &rest.Settings{Port: port})
	h, _ := trg.NewHandler(&rest.HandlerSettings{Method: "POST", Path: "/gaussian"})
	h.NewAction(RunActivities)

	// making a map of al activities to use, in this case the Inference and log activities
	//store in map to avoid activity instance recreation
	infAct, _ := api.NewActivity(&inference.Activity{})
	logAct, _ := api.NewActivity(&log.Activity{})
	activities = map[string]activity.Activity{"inference": infAct,"log":logAct}

	return app
}

var activities map[string]activity.Activity

func RunActivities(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {

	// logging
	msg:="Making an inference"
	_, err := api.EvalActivity(activities["log"], &log.Input{Message: msg})

	// Basic inputs for the inference activity
	modelPath := "Archive_simpleCNN.zip"
	framework := "Tensorflow"

	// features is the array that will be used to hold the input feature set that we pass into the inference activity
	var features []interface{}

	inp:=inputs["content"].(map[string]interface{})

	// So I read in the list of gaussin arrays from the POST request, but I need to convert them to the form the CNN model expects
	//    (for an array of shape [x,10] to a tensor of shape [x,10,1,1] where x is the number of gaussians sent),
	//    while the inference activity is taking them in as []interfaces{} at each layer
	d := inp["Input"].([]interface{})
	fmt.Println(d)
	var datafeat []interface{} 
	for _, row := range d {      // looping over each gaussian sent
		row2 := row.([]interface{})
		var gaus []interface{} 
		for _, item := range row2 { // looping over each item in the gaussian
			var it []interface{}  //Adding one more dimension
			var barebracket []interface{} // and again
			barebracket=append(barebracket,float32(item.(float64)))
			it=append(it,barebracket)
			gaus = append(gaus, it)
		}
		datafeat = append(datafeat, gaus)
	}

	// Now append the input feature with the name X to the features array. This will be passed into
	// the inference activity.
	features = append(features, map[string]interface{}{
		"name": "X",
		"data": datafeat,
	})

	// Running Inference Activity with appropriate inputs
	out, err := api.EvalActivity(
		activities["inference"], &inference.Input{
			Model:modelPath,
			Features:features,
			Framework:framework, 
			SigDefName:"serving_default",
			Tag:"serve",
			},
		)
	if err != nil {
		return nil, err
	}

	mapPred := out["result"].(map[string]interface{})

	response := make(map[string]interface{})
	response["output"] = mapPred

	//logging
	str,_:=coerce.ToString(mapPred)
	msg2:="Inference finished, output:"+str
	_, err = api.EvalActivity(activities["log"], &log.Input{Message: msg2})

	reply := &rest.Reply{Code: 200, Data: response}
	return reply.ToMap(), nil
}
