---
title: Example of Streaming Outlier Detector
weight: 4618
---

# Application
This app is an example on how to use flogo streams to handle streaming ML inferencing.  A simulation trigger that reads from csv files are used to simulate IoT sensor readings.

## Building app
### Kaggle dataset

The dataset used in this example is located [here.](https://www.kaggle.com/mlg-ulb/creditcardfraud/data)  This is a real credit card transaction dataset with transactions over a 48 hour period. there are 28 features that have be reduced for more features to anonymize the data.  Additionally, there is a time feature (0-47) and an amount feature.  A good discussion of the dataset with an analysis of its properties is [here](https://www.datascience.com/blog/fraud-detection-with-tensorflow).  The autoencoder outlier detector in the extra section of the Jupyter Notebooks is based off of this example.

### Jupyter Notebooks
Two Jupyter notebooks exist in this repo.  The first ([processing_cc_fraud_data.ipynb](processing_cc_fraud_data.ipynb)) processes the kaggle cc data and creates normalized training data and normalized sample data to run the outlier app on that the model has not previously seen.  [dnn_outlier.ipynb](dnn_outlier.ipynb) on the other hand trains the dnn model used in the app.

After running the DNN file, you must zip and move the model in the m directory into the directory that contains your binary file.

#### Additional Models
The DNN model included above is a basic classifier with a heavily unbalanced dataset.  This works great if there is labeled data available.  However, sometimes unsupervised outlier/anomaly is needed.  In the [autoencoder_outlier.ipynb](autoencoder_outlier.ipynb) notebook we build and export an autoencoder anomaly detector based upon the one [here](https://www.datascience.com/blog/fraud-detection-with-tensorflow).  The exported model (after being zipped) can be substituted into the app simply by changing the Archive file name in the flogo.json file.

### Flogo CLI
#### Building and Running the app
```
flogo create -f flogo.json outlier_app
cd outlier_app/
flogo update github.com/project-flogo/core
flogo install github.com/project-flogo/stream@master
flogo install github.com/project-flogo/stream/activity/aggregate@master
flogo install github.com/project-flogo/ml/activity/inference@master
flogo install github.com/skothari-tibco/csvtimer@master
flogo build
cd bin
```
Then the model and the data files (zero.csv, one.csv, and two.csv) from the jupyter notebooks needs to be placed in the bin directory adjacent to the binary file.  Then the app can be run from the outlier_app/bin directory:
```
./outlier
```
#### Database

This app requires that a PostGreSQL database exist on the same machine that the app runs on with a connection string as defined in the flogo.json file.  The database must have a table compatible with the outlierdatalogging custom (to this app) activity.  The table required is test.outlier and columns ind,act,pred,t. To match the Insert statement:
```
INSERT INTO test.outlier (ind, act, pred, t) VALUES ($1, $2, $3, $4)
```
This database can be used to drive an external dashboard run on Grafana (among other possibilities).

#### Debugging
Potentially, errors can arise based on old versions of modules.  Here are some commands to try in the outlier_app directory before building the app.
```
flogo update github.com/project-flogo/core
flogo install github.com/project-flogo/stream@master
flogo install github.com/project-flogo/stream/activity/aggregate@master
flogo install github.com/project-flogo/ml/activity/inference@master
flogo install github.com/skothari-tibco/csvtimer@master
```

Also, you could try updating the cli.  Outside the app directory run:
```
go get -u github.com/project-flogo/cli/...
```

