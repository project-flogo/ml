# COCO Image Recognition with Flogo

## Overview

In this example we saught to build a REST API that when sent an image file (png,jpeg,or bmp) it would run image recognition on the image and return the result.  The flogo.json builds this appn with the flogo cli.  Assuming the zip file discussed in the model section  is in the same directory as the flogo.json then run the following commands:

```
flogo create -f flogo.json
cd ImageClassification
flogo build
cd bin
cp ../../Archive_rcnn_inception_v2_coco.zip .
./ImageClassification
```

## Getting the model 

The image recognition model used in this examples was downloaded from the [TensorFlow Detection Zoo](https://github.com/tensorflow/models/blob/master/research/object_detection/g3doc/detection_model_zoo.md).  Specifically, the [mask_rcnn_inception_v2_coco](http://download.tensorflow.org/models/object_detection/mask_rcnn_inception_v2_coco_2018_01_28.tar.gz) model from the COCO-trained model section was used.  Once the tarball has been downloaded the savedmodel version of the model needs to be extracted and zipped for use in flogo.  The following commands accomplish this task.

```
tar xvfz mask_rcnn_inception_v2_coco_2018_01_28.tar.gz
cd mask_rcnn_inception_v2_coco_2018_01_28/
cd saved_model/
zip -r Archive_rcnn_inception_v2_coco.zip saved_model.pb variables/
mv Archive_rcnn_inception_v2_coco /path/to/app/bin/directory/
```

### COCO dataset
The [COCO](http://cocodataset.org/#home) dataset is a commonly used curated collection of images used in image recognition.  COCO stands for "Common Object in Context" and consists of images with backgrounds consistent with the objects detected.  The labels for the COCO dataset is listed [here](https://github.com/nightrome/cocostuff/blob/master/labels.txt), though only labels 0-91 are included in the model used here (because when the model was train those were the labels available).

### Inception models

Here we used an Inception v2 for our model because of those in the [TensorFlow Detection Zoo](https://github.com/tensorflow/models/blob/master/research/object_detection/g3doc/detection_model_zoo.md) had reasonable speed and size properties.  However, for more depth on the Inception model [here](https://towardsdatascience.com/a-simple-guide-to-the-versions-of-the-inception-network-7fc52b863202) is a simple guide to the various inception model version.