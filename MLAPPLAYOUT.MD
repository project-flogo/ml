# ML App Layout
 While there are many possible ways to use machine learning (ML) within a Flogo app, there exists a fairly standard way to create an ML app within Flogo.  Here we are going to discuss the basic layout of such an app with a few typical variants based upon triggers / usecases.

## General structure

Like Flogo in general, ML based Flogo apps are based around the consept of an event.  I.e. an event occurs, which generates data.  This data is then processed into a usable form (possibly joined with external data) so that it can be fed into the ML algorithm to make a inference.  This inference is then post-processed to allow the app perform an action.    Therefore, a typical ML app follows the following steps:
1. A Flogo trigger responds to an event.
2. Data from the trigger is processed in an activity (or multiple activities).
    * Additionally, it is possible to use the Flogo Streams to aggregate accross multiple events.
3. The processed data is input into an ML model using the inference [activity](https://github.com/project-flogo/ml/tree/master/activity/inference). 
4. The inference output is post-processed (usually based on the probability of an event or for mapping numerical category to a label).
5. A logic activity or Flogo Rules action performs some logic on the post-processed result.  It is worth noting that this logic can take place outside the Flogo App (say in a web interface).
6. A final activity performs an action whether that is returning a response for a REST trigger or sending down a shutdown command to an industrial IoT device with an MQTT message activity.

It is also worth noting that an activity to log or export data to a database can occur at any step.

## Examples

Examples of this layout exist [here](https://github.com/project-flogo/ml/tree/master/examples).