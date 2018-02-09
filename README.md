# Welcome to Go Hackathon Starter!

## Why did I create this? 
When I first started out with Go, I found tons of comments online saying - just use the standard library! Now Go's standard library is one of the best one's I've seen, but it doesn't give all you need to build a full fledged web application. I tried to use go for many of side projects and found myself repeating some of the most common things like authentication, user management, lack of generators for CRUD applications etc.
This project is heavily inspired by [hackathon starter](https://github.com/sahat/hackathon-starter) which is based on nodejs and express. Being a developer with both nodejs and Go experience, I found this missing in the Go ecosystem. This project is a step towards that direction. The focus of the project is to optimize developer speed and keep everything clean and reusable as much as possible.

**Live Demo**: TODO host this

<!-- TODO: add screenshots -->

Working in Progress
-------------------
**TODO:**
- fix UI on login page
- add setup.sh to install replace imports with your directory names
- Update Readme for better instructions to setup integrations for login and api-examples
- add tests and CI
- deploy working example to heroku with all auths working

- add account linking (needs more design)
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



Table of Contents
-----------------

- [Features](#features)
- [Prerequisites](#prerequisites) 
- [Getting Started](#getting-started)
- [Obtaining API Keys](#obtaining-api-keys)
- [Project Structure](#project-structure)
- [Recommended Go Libraries](#recommended-go-libraries)
<!-- - [Recommended Client-side Libraries](#recommended-client-side-libraries) -->
- [FAQ](#faq)
<!-- - [How It Works](#how-it-works-mini-guides) -->
- [Cheatsheets](#cheatsheets)
    <!-- - [Adding new authentication](#) -->
- [Benchmarks](#benchmarks)
- [Deployment](#deployment) TODO
- [Docker](#docker) TODO
<!-- - [Changelog](#changelog) -->
<!-- - [Contributing](#contributing) -->
<!-- - [License](#license) -->


Getting Started
---------------

You will need to clone the project and run following commands - 

```
# cd into project directory
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

Obtaining API Keys
------------------
**Github**
**Facebook**
**Twitter**
**Google**

Project Structure
-----------------
Project follows MVC architecture


Recommended Go Libraries
------------------------
- gobuffalo
- goth
- testify
- delve

FAQ
----

**Unable to login/ start fails with .env file not found error**
- You will have to set keys for github, facebook, twitter to get the login working for you.
- will be adding detailed instructions in Readme itself in next few days

**New project build fails due to import failure**
- this is because imports were not updated are are still using `github.com/niranjan92/go-hackathon-starter`
- you should run `setup.sh` script that will find and replace all the dependencies with correct paths


CheatSheets
-----------
**Adding new authentication**

Benchmarks
---------------

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

Docker
---------

Deploy
----------

[Powered by Buffalo](http://gobuffalo.io)


