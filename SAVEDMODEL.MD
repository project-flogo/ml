# SavedModel format

The supported format for the ML [inference](activities/inference) activity is the TensorFlow [SavedModel format](https://www.tensorflow.org/guide/saved_model).  This format consists of two parts.
1. saved_model.pb -  a protobuf file contains a serialized representation of the graph of the model
2. variables -  a directory that contains the values of all the variables in the nodes in the graph.