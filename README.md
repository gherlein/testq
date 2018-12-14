# Test Question

## Background

This is a homework assignment I have been known to give to candidates.  Here's what I tell them:

```
Please follow these instructions exactly.  Please write a little program in the language of your choice to query the http://www.omdbapi.com/ OMDB API.  
Here’s the requirements:
 
Command line only – no GUI wanted or needed
Must run on linux
Any language you choose
Program must accept a command line parameter of a name of a movie
Program must query the API for a result
Program must output the Rotten Tomato rating in some useful way to the user
Program must be “dockerized” – i.e. the work to make it into a docker container must be present
All source code for the program and dockerization must be in a git repository
The Git repository must be in your github.com account and publicly accessible (github.com is free for public repos)
You should assume I will clone your repo and test it personally
```
## Using the Required API Key

The [OMDb API](http://www.omdbapi.com/) requires that you get an API key.  It's free for up to 1000 queries a day which is fine for our homework assignment.  This program requires that you set that key as an ENV variable, like so:

```
export OMDAPI_KEY="XXXXXXXX"
```

This is STILL NOT SECURE but it's a slight improvement over passing on the command line.  Command lines are visible in process space and thus easily sniffed.  

In this case, it's all a moot point anyway since the API provider does not use SSL so HTTPS is not an option.  This enables anyone sniffing the network to pluck the key right out of the query and getting access to the network is far easier than getting access to the host this would run on.  But in the interest of good coding hygiene it's in an ENV variable.


## Usage

To build, simply type 'make' and it will build.  To execute it provide a single argument that is the name of the movie.  Example:

```
./testq -title="Hunt"
The Hunt: 94%
```

You can also specify a debug mode to see the actual query it makes to the API:

```
./testq -title="Hunt" -debug=true
http://www.omdbapi.com/?apikey=XXXXXXXXXX="Hunt"
The Hunt: 94%
```

WARNING:  debug mode will expose your API key in the query.

## Docker Container

By default the Makefile will also build this as a docker container.  You can exectute it like this:

```
docker run -ti -e OMDAPI_KEY gherlein/testq -title="Hunt"
The Hunt: 94%
```

# License

This project is released under the MIT License.  Please see details [here] (https://gherlein.mit-license.org).
