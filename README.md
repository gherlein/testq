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
## Usage

To build, simply type 'make' and it will build.  To execute it provide a single argument that is the name of the movie.  Example:

```
testq gherlein$ ./testq "The Shawshank Redemption"
http://www.omdbapi.com/?apikey=6a7638ee&t="The Shawshank Redemption"
```



# License

This project is released under the MIT License.  Please see details [here] (https://gherlein.mit-license.org).
