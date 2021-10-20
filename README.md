# paddington
WholeNumPad takes a string and an integer, and will return a string
with each iteration of whole numbers found within the string front
padded with zeros.  ex: fn("James Bond 7", 3) ==> "James Bond 007"

## Madness
Once your learn recursive, you can't help see recursive solutions.
They are not always the way to go, and often in non functional languages
they are more a pain and complexity not worth the time... in this case
the problem was simple enough, and I thought to use Go functionally.

## Test it
```shell
  go test -v
```