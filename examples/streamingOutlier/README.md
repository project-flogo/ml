---
title: Example of Streaming Outlier Detector
weight: 4618
---

# Application
This app is an example on how to use flogo streams to handle streaming ML inferencing.  A simulation trigger that reads from csv files are used to simulate IoT sensor readings.

## Building app
### Kaggle dataset


### Jupyter Notebooks
Two Jupyter notebooks exist in this repo.  The first (processing_cc_fraud_data.ipynb) processes the kaggle cc data and creates normalized training data and normalized sample data to run the outlier app on that the model has not previously seen.  dnn_outlier.ipynb on the other hand trains the dnn model used in the app.

After running teh DNN file, you must zip and move the model in the m directory into the directory that contains your binary file.

### Flogo CLI
```
flogo create -f flogo.json
cd outlier_app
flogo build
```
Then the model and the data files (zero.csv, one.csv, and two.csv) from the jupyter notebooks needs to be placed in the bin directory adjacent to the binary file.

