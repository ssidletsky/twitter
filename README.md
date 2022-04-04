# Task

1. Write a database schema for creating the tables in a MariaDB/MySQL database with users, tweets, and followers.

* A user has a username, password, email, first name, last name, and age.
* A tweet has text, publication date, and is posted by a user.
* A user can be followed by other users.

The tables may contain additional structures that provide better space efficiency, security, or performance.

2. Write REST API for the tweet. It should be able to return the tweets by users followed by a specific user. The result must include username, first name, last name, tweet text, and publication date and support pagination.

3. Write a utility function for a User object. You will be implementing three different functions, each with the following function signature: func ([]User) bool.

   * AtLeastTwice is a function that returns true if there is a person who is at least twice as old as any other person in the list, otherwise, the function returns false.
   * ExactlyTwice is a function that returns true if there is a person who is exactly twice as old as any other person in the list, otherwise the function returns false.
   * ConstrainedExactlyTwice is a function that behaves like ExactlyTwice, but input age values are guaranteed to always be within the range 18 to 80, and this function must perform very well (consider time and space complexity).

# Solution

## Database

The database schema can be found [here](db/init.sql).

## Rest API

Rest API and MySQL are configured to run inside docker containers. If the MySQL container starts for the first time, the schema and some initial data for testing will be created automatically. These SQL scripts can be found inside `db` folder.

Configuration files for API are located inside `config` folder. There are two configs to run the API server locally or inside docker.

### API architecture. 

I tried to follow Clean Architecture principles while designing the app. As a result, the API isn't dependent on MySQL even though it's a requirement in the task. All the communication with DB is located at [repository layer](app/tweets/repository/mysql/tweets.go). The real dependency is an [interface](app/tweets/repository/tweets_querier.go), but `mysql` repository is only implementation of it. As a result, I can switch to the different databases only by implementing the required functionality defined by the corresponding interface. These ideas can be moved further to apply to other dependencies like API framework (fiber) or to logger package (logrus).

There are three more layers: entities, use cases, and delivery. Entities are business objects of the application that encapsulate business rules and can be either object with methods, or they can be a set of data structures and functions. Use cases are a layer that contains application-specific business rules. A delivery layer is responsible for how the application is being delivered. In this case, it's via HTTP, but could also be gRPC or CMD.

### Implementation details

Entry point to the API is [api.go](api.go). The script configures and runs the API server.

The pagination is done by specifying the number of tweets per page by `per_page` parameter and the latest seen tweet ID  by `from_tweet_id`. I purposefully didn't use MySQL's `OFFSET` and `LIMIT` under the hood, as the performance will be decreased for huge amounts of data and a big number of `OFFSET`.

## Utility functions for User object

All the functions are implemented at [app/tweets/entities/user.go](app/tweets/entities/user.go).

I couldn't find any advantages of knowing the range of ages to implement a better version of the `ExactlyTwice` function. Even though I tried to not access the map in the cases it doesn't make sense. For example, there is no sense to look up users twice as old as the user of 40+ age. The same for the uses with age <36, there is no sense to look up the users twice as younger. Even with these improvements, the performance of the func didn't change.

I left the improved implementation, but most likely I would use `ExactlyTwice` instead of it.

# Run

In order to start dockers run:

```bash
make up
```

In order to request tweets for user with id `1`:

```bash
curl -H "X-USER-ID: 1" localhost:8000/tweets\?from_tweet_id=0\&per_page=10
```

###### Response:
```bash
{"tweets":[{"id":8,"username":"joey","first_name":"Joey","last_name":"Tribbiani","text":"Why do you have to break up with her? Be a man. Just stop calling.","publication_date":"2022-04-04T15:53:08Z"},{"id":7,"username":"joey","first_name":"Joey","last_name":"Tribbiani","text":"Joey doesnâ€™t share food!","publication_date":"2022-04-04T15:53:08Z"},{"id":6,"username":"phoebe","first_name":"Phoebe","last_name":"Buffay","text":"Oh My God, A Woman Flirting With A Single Man? We Must Alert The Church Elders!","publication_date":"2022-04-04T15:53:08Z"},{"id":5,"username":"phoebe","first_name":"Phoebe","last_name":"Buffay","text":"If You Want To Receive Emails About My Upcoming Shows, Then Please Give Me Money So I Can Buy A Computer.","publication_date":"2022-04-04T15:53:08Z"},{"id":4,"username":"monica","first_name":"Monica","last_name":"Geller","text":"Theyâ€™re as different as night and... later that night.","publication_date":"2022-04-04T15:53:08Z"},{"id":3,"username":"monica","first_name":"Monica","last_name":"Geller","text":"Having a heart attack is natureâ€™s way of telling you to slow down.","publication_date":"2022-04-04T15:53:08Z"}]}
```
