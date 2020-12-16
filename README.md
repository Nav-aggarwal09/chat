# ASAPP Chat Backend Challenge v1
## Overview
This is a Golang implementation of the HTTP Server defind in [ASAPP's coding challenge](https://backend-challenge.asapp.engineering/).

This Readme serves to provide a high level overview of the implementation, challenges, and Wishes/ToDos while going about this.
We will go package by package..

#### Requirements completed
The `/check`, `/users`, and `/login` are implemented successfully run. A SQLite3 database was also created. 

### Cmd

Starting with the entry point, `cmd/server.go` creates a Users table in the root as well as setting 
a logging framework to be used throughout the project. The http handler function were given as backbone 
and needed to further changes. 

##### Possible Cmd Enhancements
- take in more command line arguments 
    - publish a different port
- Have Users's table timestamp column be registered as a type DATE/TIME to make future queries easier
- Created more tables to be used for sending/receiving messages

### Auth
Auth's only function is to validate the user. The vision and steps in mind are in the code, but to repeat them,
`ValidateUser()` validates a token by:
1. fetching sender's token from the DB and comparing with given token
2. If yes, check if the token is expired (older than three hours from timestamp)
3. If not old, do nothing(?)
4. If old or token not matched, return StatusUnauthorized

#### Auth Challenges
A big question for me was step 3. If everything was ok, how do I proceed to the send/get message functions? 
ValidateUser returned a `http.HandlerFunc`, but how does that translate to hitting the other functions since 
we are basically a function within a function. 

Another step I felt clunky about was getting the token and senderID. I needed both to ensure that the right 
sender was sending the right token. However, the former ID was in the body where I had to parse it out of the string
and the latter was in the header which was easily taken out.  

### Controllers
The login, users , and message controllers just translate the request to a string and pass it off to 
the model for it to take care of. I did not want to handle further processing of data as I believe that is the 
models job. 

Health controller just returned a models.health struct when hit. 

#### Controllers Wishes
For health, its pretty bare bones. I guess I would call the functions internally or run some tests internally to 
ensure that everything isn't returning errors.

Otherwise it all was pretty straightforward. Get the data from the request, turn it into something the model
can handle, and ship it to the model.

### Database
This existed to eliminate cyclic dependency that Golang doesn't allow. 

#### Database Wishes
Ideally, I would have like to create functions to handle the quries being run throughout the program.

### Models
User.go handles everything users related -- creating/storing profiles and authenticating

Health.go contains the model struct

login.go is being used to generate token and timestamp for the DB when logging in. 

message.go is storing and fetching (still to be implemented) messages

#### Model Wishes and Challenges
Users.go wasn't that bad. The only caveat is that I have not put any encryption when storing 
passwords. 

Messages.go was weird since there could be 3 types of messages. I categorized them into a `Content interface`. 
What I was going to do for the DB was store them into a big table with the specified columns:
* senderID
* receiverID
* timestamp
* MsgType
* text
* Url
* Height
* Width
* Source
For every Msg type, I would fill out their respective fields and leave the rest nil. 
   
   
## Takeaways
Overall fun, open-ended project. Definetley learned a lot. Pretty happy with where I got with my limited knowledge 
as an intern and limited time.   


