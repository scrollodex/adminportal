# scrollodex -- The scrollodex CMS

Scrollodex is a web-based system for maintaining a static website that lists professionals (doctors, plumbers, therapists,
etc.).

The raw data is stored in YAML files in a git repository (rather
than using a database system like MySQL).

An interactive web-based "admin portal" permits easy editing and
updates.

This data is used to generate input to a Hugo-based static website.


https://medium.com/@saif_nasir/connect-docker-containers-the-easy-way-60fae730fe0
sh exec.sh
docker run -p 6379:6379 --name some-redis -d redis redis-server --save 60 1 --loglevel warning

https://stackoverflow.com/questions/42385977/accessing-a-docker-container-from-another-container

https://docs.docker.com/compose/networking/

Starts all containers:
start-all.sh

docker exec -it myportal bash
apt-get update && apt-get install -y netcat


# How to re-clone the test repos.
REPO=bi
cd ~/gitthings/test-db-$REPO
rsync --delete --exclude=.git -avP ../scrollodex-db-$REPO/. .
git add README.md category entry location 
git commit -m'clone from source' -a
git push

# The admin panel:
  https://beta.scrollodex.net
  https://www.scrollodex.net


# Creating users

Have them log in and be rejected. Give the rejection code:

Add code here:
vi routes/unauthorized/unauthorized.go


# Architecture

CURRENT: (phase 1)

Webserver + MySQL
-> hourly check; generate yaml commit to hugo-bi
-> CircleCI sees change in hugo-bi, generates website

NEW: (phase 2)
-> adminportal edits scrollodex-db-bi
-> clicks "publish": generates yaml and commits to hugo-bi
-> CircleCI sees change in hugo-bi, generates website

# Re-clone the beta area from production:

make reset-test-data

#

# Website build process:

Phase 1:
Hourly cron job on JJ reads SQL and creates all the individual files.

* Runs p1extract to create the the individual files and the single YAML file:
Individual entries: Stored in git:scrollodex-db-{bi,poly}:/{category,entry,location}
Single yaml file: Stored in git:hugo-{bi,poly}:/data/entries.html

* CI/CD pipeline sees the change in hugo-{bi,poly} and regenerates the
  websites.

Phase 2 BETA:

* adminportal "edits" the files in git:scrollodex-db-{bi,poly}
* 

Phase 2:

1. Disable cron job
