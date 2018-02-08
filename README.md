# Welcome to Go Hackathon Starter!

## Motivation
A common complaint for folks starting out with golang is no availabitliy of commonly available. This project is heavily inspired by [hackathon starter](https://github.com/sahat/hackathon-starter) which is based on nodejs and express. Being a developer with both nodejs and Go experience, I found this missing in the Golang ecosystem. This project is a first step towards that direction.

## Getting Started

### Setting up the project

You will need to clone the project and run following commands - 

```
# cd into your home directory
cd ~/go/src/github.com/<username>/

# clone the code
git clone git@github.com:niranjan92/go-hackathon-starter

# install dependencies
npm install

## find and replace imports with your directory
## eg. import "github.com/niranjan92/go-hackathon-starter" needs to be replaced with
## 	   import "github.com/<your_username>/go-hackathon-starter"
## this will soon be fixed with new utility script `setup.sh` - coming soon

# init go dependencies
dep ensure

# The first thing you need to do is open up the "database.yml" file and edit it 
# to use the correct usernames, passwords, hosts, etc... that are appropriate for 
# your environment. then create database
buffalo db create -a

# run migrations
buffalo db migrate

# start server in dev mode
buffalo dev

```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Go Hackathon Starter!" page.

## Benchmarks
Go is pretty fast and focus should be on adding new features quickly. But for those of you looking for benchmarks -
Tested with `wrk` on home page and was able to run at ~370 requests per second with average latency of ~3 ms. This way more than
what a budding project needs. :)
```
Niranjans-MacBook-Air:hackathon-starter niranjan$ ~/go/bin/wrk -t1 -c1 -d20s http://localhost:3000
Running 20s test @ http://localhost:3000
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.02ms    2.86ms  48.97ms   92.86%
    Req/Sec   375.58     97.42   554.00     65.00%
  7493 requests in 20.05s, 33.95MB read
Requests/sec:    373.73
Transfer/sec:      1.69MB
```

## FAQ

### Unable to login/ start fails with 
- Check if `.env` file is present in the project directory
- You will have to set keys for github, facebook, twitter to get the login working for you.
- will be adding detailed instructions in Readme itself in next few days

### new project build fails due to import failure
- this is due to incorrect 

**TODO:**
- Update home page with details
- add tests and CI
- add setup.sh to install replace imports with your directory names
- Update Readme for better instructions to setup integrations for login and api-examples

- FACEBOOK - show list of friends
- add profile pic / gravatar to profile age

**Good to do:**
- delete account feature
- add paypal test example
- fix highlighted part for selected tab in nav
- admin management console for users
- add a simpler way to configure db. maybe using docker compose

**Maybe:**
- add examples for angular and react

[Powered by Buffalo](http://gobuffalo.io)


