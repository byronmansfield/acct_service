# Account Service proof of concept

This is a proof of concept restful go server that hopefully will replace my works current webservice. Also it is used as a project I for me to learn golang.

## Usage

Make sure you have postgres running. Right now being this is a proof of concept the app expects to have a database todos with a table called todos. If you do not have those you will want to add them to your local postgres first.

Next you want to cd into the project and build and run the executable like so:

```
go build && ./acct_service
```

Now from another shell you can call the todos endpoint with curl to get all the todos like this:

```
curl -i localhost:8080/todos
```

You can post to it like this:

```
curl -X POST -H "Content-type: application/json" -d '{"id":"4","name":"go to store","completed":"false","due":"2015-09-08"}' localhost:8080/todos
```

Thats it so far. I still need to refine some of these steps and make them a little more graceful. Also I think I still need to do the get by id. Along with the rest of the roadmap. But its up and running with very little effort or code. Super fast and lightweight.

## Roadmap

- [ ] Concurrent Server
- [ ] CORS support
- [x] Postgres connection
- [ ] Redis
- [ ] API versioning
- [ ] RPC support 
- [ ] Auth support 
- [ ] Swagger integration

