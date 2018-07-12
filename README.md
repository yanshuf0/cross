# cross
golang assessment

This is a go server with api endpoints that will return data about products for a coffee company.

## api
`/api/product/machine` - accepts a `size_id` query_parameter, if left blank all coffee machines will be returned.
`/api/product/pod` - accepts a `flavor_id` and/or a `size_id` query parameter. Left blank will return all pods.
`/api/cross/machine` - requires a `coffee_machine_id` query parameter. Will return the smallest packs of pods, one per flavor.

`/api/cross/pod` - requires a `pod_id` will return all other pods with a matching `flavor_id` and `size_id` (excluding itself).

## frontend
There is an angular application I mocked up to that gives a (crude) visual representation of most of the api features. 
It is located in /web/cross-spa. To build it cd to that directory and run `ng build --prod`. 
This outputs the dist directory to `/assets`

## server
The server can be built if you cd into the `cmd` directory and run `go build`.

## flags
The program accepts two command line flags `assetsDir` which is the path to where assets are located 
(the text file with the data, the .sql file acting as a serverless db and the spa files. The second argument is `port`
where you can specify the port to run the program on. Both flags have defaults and do not have to be specified 
during development.

## data
I've created a sqlite3 relational database that reads the product descriptions from a plain text file and converts the data
to the relational model. It's created from scratch on each run to preserve integrity for assessment purposes only.
![database scheme diagram](https://raw.githubusercontent.com/yanshuf0/cross/master/db-schema.png)

## use
I've uploaded the application to my server feel free to use the front end application to test the endpoints or use postman
or your tool of choice. Let me know if you have any questions.

The url is:
http://35.236.46.43
