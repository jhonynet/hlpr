# HELPER

# usage
hlpr workflow{.json,.yaml}

## use cases:
* dump file to an api.
* dump api to a file.
* map csv fields to an api call and then save the resultant object as json stream file.
* build object with multiples y/o parallel api calls and then post it to an api.
* run an http server and use each request as part of a workflow.

## TODO:
* workflow validation before run.
* http better body parsing.
* custom template function to use global vars.