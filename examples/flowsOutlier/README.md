# Flows Outlier 
## Overview
This example takes the [streamingOutlier](../streamingOutlier) example and joins the csv files together into a single file that is read as a flow.  The rest of the description is the same.  The [cc_outlier.csv](cc_outlier.csv) and [Archive_20190315.zip](Archive_20190315.zip) files were generated using the Jupyter notebooks from [streamingOutlier](../streamingOutlier).  The one addition is instead of installing streams we need to install flows (and we still need to update the inference activity and the csv trigger):
```
flogo create -f flogo.json
flogo install github.com/project-flogo/flow
flogo install github.com/project-flogo/ml/activity/inference@master
flogo install github.com/skothari-tibco/csvtimer@master
flogo build
cd bin
./flow_outlier_app
```