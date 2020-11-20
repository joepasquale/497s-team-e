# Scheduler Service Documentation

## Service Author
Nick Thurai

## Description
For scheduling we are running a python script which will take in a json object representation of a group and the group’s availability and from there will look for available times among the group. We also included multiple different formats to make it easier during testing for the group to use the scheduler. For scalability concerns this script is not resource intensive and can be used to fit any number of individuals or requests. The internal functionality takes in the json object and iterates through to find the top times that best work. This function can also be used in the testing phase. 
