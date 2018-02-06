TASKS:
- FACEBOOK - show list of friends (and their birthdays??)
- add sample react and Angular pages
- add tests
- linked accounts explore in detail
- add profile pic / gravatar link to my account dropdown and profile page

GOOD to do
- add paypal test example
- save account profiles
- delete account feature?
- fix highlighted part for selected tab in nav

- admin management console for users ??? (in scope for hackathon starter? maybe no..)

WAY_INTO_FUTURE:
- later add pluggable support for angular and react

# Welcome to Go Hackathon Starter!

# TODO: add a simpler way to configure db. maybe using docker container, or check if docker-compose is supported???
## Database Setup

It looks like you chose to set up your application using a postgres database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a
## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Go Hackathon Starter!" page.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)


