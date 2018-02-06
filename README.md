# Welcome to Go Hackathon Starter!

## Getting Started
### Database Setup

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

### Create Your Databases
Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a
### Starting the Application
	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Go Hackathon Starter!" page.

TODO:
- FACEBOOK - show list of friends (and their birthdays??)
- add tests
- add profile pic / gravatar to profile page
- delete account feature
- add a simpler way to configure db. maybe using docker compose

GOOD to do
- add paypal test example
- fix highlighted part for selected tab in nav

- admin management console for users ??? (in scope for hackathon starter? maybe no..)

MAYBE: 
- add examples for angular and react

[Powered by Buffalo](http://gobuffalo.io)


