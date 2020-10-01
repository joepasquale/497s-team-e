# Accessing Yelp API

## Group E

This is a quick example of how to use python to communicate with the Yelp APIs which we will use for our website to return data on potential destinations for groups to meet up.

As of now, all you need to do is send a POST request to https://ip:5000/yelp with the parameters of terms and location. Terms is the contenxt of the business you're looking for. Location is the location you want to get the list of businesses in that area. There are other optional parameters which can be added later, one current one being used is the limit which is currently hard coded to a max of 10 businesses being shown, but easily can be added as another parameter later.