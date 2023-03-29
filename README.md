# qdb

**qdb** is a minimal key-value storage based on SQLite that has a dead-simple convention for using it.

## Usage

- Run `go run .` to start the server locally at `http://localhost:3000`.
- If this is your first-run, note the admin password that is printed in console.
- Call the API using this guide (assuming the `admin` password is `secret`):

### Set a new value

Replace `KEY` and `VALUE` with your own key/value:

```sh
curl -X POST http://localhost:3000/KEY -d 'VALUE' -u 'admin:secret'
```

### Retrieve a value

Replace `KEY` with your own key:

```sh
curl http://localhost:3000/KEY -u 'admin:secret'
```

### Delete a key

Replace `KEY` with your own key:

```sh
curl -X DELETE http://localhost:3000/KEY -u 'admin:secret'
```

### Create a new user

Replace `USERNAME` and `PASSWORD` with your own username/password:

```sh
curl -X POST http://localhost:3000/users/USERNAME -d 'PASSWORD' -u 'admin:secret'
```

### Delete a user

Replace `USERNAME` with your own key (you can't remove the `admin` user):

```sh
curl -X DELETE http://localhost:3000/users/USERNAME -u 'admin:secret'
```
