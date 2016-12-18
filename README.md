### Convenience utility
Loads CSV in the specific form into redis hash using redis protocol.
Inspiration https://github.com/tlehman/redis-mass

### spec
MIME "text/csv" https://tools.ietf.org/html/rfc4180
  
```
col1, col2
field, value
```
e.g.

File: users.csv
```
1, Alice
2, Bob
```
is equivalent to:
```
HMSET users 1 Alice 2 Bob
```
If there is a comma ',' in the file name it is replaced with colon ':'. So:
```
users,1.csv
```
Will end up in the key:
```
users:1
```
Example use:
```
hmset-csv users.csv | nc localhost 1111
```
Content of users.csv will be piped to the localhost:1111 hash key 'users'. 'nc' - is netcat


### download binary
  https://github.com/biosckon/hsmset-csv/releases/tag/v1.0
### install
  copy to one of directories in your PATH e.g. /usr/local/bin or C:\Windows\
  
 
