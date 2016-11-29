### Convenience utility
Loads CSV in the specific form into redis using redis protocol.
- using some code from https://github.com/tlehman/redis-mass
```
col1, col2, col3
key,  field, value

e.g.

users,   ,
     , 1 , Alice
     , 2 , Bob
```
is equivalent to:
```
HMSET users 1 Alice 2 Bob
```

