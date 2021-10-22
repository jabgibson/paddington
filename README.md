# paddington
### update [10-22-2021]
The first version of this had only one solution (see earlier commits). Because I was unsure if the party
requesting me to write this would have any concerns about the recursive nature of the first solution
I have now created two solutions. Both were from the hip, without much research or diligence in presentation.

The test file... shared_test.go is merely holding the test parameters that both tests will run.

### original introduction
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