# Machine Learning in Project-Flogo

## Overview

As the name suggests, this project-flogo/ml repo is intended to contain all the modules for Flogo that is tied to machine learning (ML).  The primary tool for performing ML in Flogo is the [inference activity](https://github.com/project-flogo/ml/tree/master/activity/inference), which is built around a framework based approach that allows for easy extension of the activity to include different ML frameworks.  Currently, the only framework that is implemented is __TensorFlow 1.X__.  Consult the[ Tensorflow in Flogo](https://tibcosoftware.github.io/flogo/development/flows/tensorflow/) documentation for an overview of requirements and inference inputs.

This repository follows the directory structure of [Flogo Contrib](https://github.com/project-flogo/contrib).  Namely, the root directory contains the [activity](https://github.com/project-flogo/ml/tree/master/activity), trigger, function, etc. directories each of which contain the correspondingly named flogo structures relevant to ML (if these structures exist for the ML repository).  Along with the directories for the flogo structures an [examples](https://github.com/project-flogo/ml/tree/master/examples) directory that contain relevant examples for implementing ML in flogo.

## Activity/Inference 
The Inference activity is built in such a way as to allow multiple ML libraries to be called.  Each library is called a framework (which corresponds to an inference input below).  However, at the moment only TensorFlow 1.X is supported as a framework leveraging the Golang API from TensorFlow.

### TensorFlow Framework

First and foremost, Flogo supports inferencing TensorFlow models, it does not support training of models using incoming data. The training should be performed in Python and the SavedModel format exported for inferencing at runtime in Flogo.

Before you can begin inferencing TensorFlow models within your Flogo Flows, you’ll need to consider a few requirements.

#### Pre-requisites

The TensorFlow dynamic lib must be installed on both your development machine, as well as the target machine/device. The dynlib must be built specifically for your platform architecture, that is, Linux Arm, x86, x64, Darwin, etc. Follow the instructions documented by [TensorFlow](https://www.tensorflow.org/install/install_go), note that the only steps that you'll need to follow are 2 and 3: downloading the correct dynamic lib and setting your lib paths. You do not need to 'go get' TensorFlow.

#### TensorFlow Models
As previously stated, Flogo is leveraged to inference models at runtime, not train any models. Flogo includes a native activity to inference models. The activity has been developed and tested against the output of the [tf.estimator](https://www.tensorflow.org/api_docs/python/tf/estimator) API from TensorFlow as well as manually built models saved with the [tf.saved_model](https://www.tensorflow.org/api_docs/python/tf/saved_model) module.

#### Inputs

##### model

The “model” input to the activity should be either of the following:

- An archive (zip) of the TensorFlow model
- A directory containing the exported protobuf and check point files

The activity has been tested with the exported model from the [tf.estimator.Exporter.export](https://www.tensorflow.org/api_docs/python/tf/estimator/Exporter) operation as well as manually built models exported with the [tf.saved_model](https://www.tensorflow.org/api_docs/python/tf/saved_model) module. After export, optionally zip the file, where the saved_model.pb file is located at the root of the archive.

The SavedModel format contained in the protobuf includes metadata (interpret this as an instruction manual) on how to use the model.  The below inputs and outputs are which parts of the metadata to use to connect to this model.

##### features

The data to be passed into the SavedModel are defined in “features”. This is an array of maps and should match the following format.  For estimators an example of “features” is:

```json
[
{
    "name": "inputs",
    "data":{
         "z-axis-q75": 4.140586,
         "y-axis-q75": 4.140586
    }
}
]
```
And, for manually build models with multiple inputs “features” would be something like:

```json
[
{
    "name": "X1",
    "data":[
        [1,2,3],
        [4,5,6],
        [7,8,9]
     ]
},
{
    "name": "X2",
    "data":[
        [0.1,0.2,0.3],
        [0.4,0.5,0.6],
        [0.7,0.8,0.9]
     ]
}
]
```

##### framework

The deep learning framework to use for inferencing. Currently the only supported value is: **Tensorflow**

##### tag (default: “serve”) and sigDefName (default: “serving_default”)

A TensorFlow model consists of a complex graph (network) of mathematical operations that can contain many moving parts.  Another way to consider this is that a model consists of connected “computers” that each have a purpose.   To use the model we have to know which computer/part of the model to use.  “tag” is used to identify within the model metadata which part of the model to use.  Once we have selected the “computer” to use we then need to know which “ports” to use.  “sigDefName” is used within the model metadata to properly connect Flogo to the model.  These Inputs into the inference activity default to the standard values used by TensorFlow.

#### Outputs

The output is an object and can contain multiple outputs. For example, for a classification model, the scores and classifications will be held in a map:

```golang
map[scores:[[0.049997408 0.010411096 0.93959147]] classes:[[Jogging Sitting Upstairs]]]
```


## Examples

Currently no examples have been ported to this version (0.9.0) of flogo.



