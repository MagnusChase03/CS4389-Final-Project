# DB

*The databases are responsible for keeping persistant data, caching data, and
pub/sub for real time chat.*

## Documentation

The persistant storage of **MariaDB** is located in ./mariadb/data. To re-init
the database, the data directory must first be deleted:

```
$ sudo rm -rf mariadb/data
```
