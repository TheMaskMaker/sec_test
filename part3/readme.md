TODO if this was a real api

- use a framework like echo or gin instead of base http
- add multi get and multi post
- header checking
- better logging
- actual security (auth middleware)
- move routing out of main
- move filename to config with viper
- get rid of the panics, those are for my convenience not for prod
- user interfaces to break up ownership levels 
- use testify instead of base testing package

TODO to put this on the cloud:
- dockerize, dual stage build for small image
- ECR->EC2 with some gateway in front, might need to mess with write perms for in the docker image to make sure I can overwirte it
- add extra protections like file size/ map size limit
- put a gateway service or reverse proxy in front of it
- though realistically I wouldn't use a file db in the first place, so swap that out fo dynamo or something
