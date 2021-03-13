# upwork-scraping
Extract jobs and profile info from Upwork logged user 

# Architecture
![architecture](https://github.com/otherpirate/upwork-scraping/blob/main/doc/architecture.png)

# Solution
The main process is:
1) Login into the Upwork platform with the user given information
2) Access the main page to extract data about jobs, and save it on the jobs store folder
3) Access contact info and profile page to get the user information relevant to Agryle.
4) Put the extracted data into a JSON model as required [here](https://argyle.com/docs/api-reference/profiles)
5) Save profile.json to store folder
6) Send a message into rabbitMQ queue to all other system knows that we have a new/updated profile 

The below process runs to every new message that came from RabbitMQ. So we can scale the upwork-scraping based on the size of that queue.

When an error occurs, even in any step of crawler, the message is gonna be requeued [10](https://github.com/otherpirate/upwork-scraping/blob/main/pkg/settings/settings.go#L34) times, if the error keeps going, it will be put in an isolated queue.

I did it because can be a problem in Upwork(maintenance for example) and we can process the message an after momentnt

Another important point is, the selenium (heavy process) is started a single time on startup of docker/pod

# Libs
* Selenium to crawling/navigation
* Soup to parser HTML and get nodes
* AMQP to send/receive messages from RabbitMQ

# Run

#### RabbitMQ: 
```
docker run --rm --hostname rabbit-scraping --name rabbit13 -p 15672:15672 -p 5672:5672 -p 25676:25676 rabbitmq:3-management
```

#### Start as bin/go:
```
cd cmd
go run upwork.go
```
#### Start as docker
```
make docker
make docker_run
```

### Trigger
Publish a message in RabbitMQ: http://localhost:15672/#/queues/%2F/upwork-scraping-user

`{"username": "bobsuperworker", "password": "Argyleawesome123!", "secret_awnser": "Bobworker"}`

You can also add a field called extra_data in profile format to send extra data about the user (i.g: ssn, gender)
```
{"username": "bobsuperworker", "password": "Argyleawesome123!", "secret_awnser": "Bobworker", "extra_data": {"gender": "male", "ssn": "12345"}}
```


### Outputs
JSON files gonna be in `/tmp/upwork-scrapping/store_json` to go run and in a docker volume called `docker_store_json` to docker run 

RabbitMQ should have a new message with the profile created/updated: http://localhost:15672/#/queues/%2F/upwork-scraping-profile


# Problems
I gotta a little confused about the profile API reference, there are a few deprecated fields, but the example still using that.

Upwork block me every time, I could just do a single login in hours, so I mocked the HTML responses to keep going with the parser

My major problem is with "I'm not a robot" page when I access the contact info, I could not get deep into why it is happening. (Maybe selenium, for chrome config, not sure).

# Future Work
I did not get the full email from user information, I think I can get it after clicking on the edit button

A first step should save user, pwd, and secret on an isolated database, to keep the relevant data storage for a long period. With it, we can refresh user information or get more data once Upwork make it available 

Increase the test coverage, I just did a full unit test, checking the profile data. The coverage is good but can be really better, checking errors and corner cases.


