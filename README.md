# WORK IN PROGRESS
  
### Convenience utility
Loads CSV in the specific form into redis hash using redis protocol.
Inspiration https://github.com/tlehman/redis-mass
```
col1, col2
field, value

e.g.

File: users.csv
1, Alice
2, Bob
```
is equivalent to:
```
HMSET users 1 Alice 2 Bob
```

