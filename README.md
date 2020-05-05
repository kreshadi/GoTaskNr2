Utility Library

Write a "small library" with utility functions

Examples:

Mathematical operations (something that is not in Go "math" package, maybe little cryptographic things)

Slice and map operations (delete item at spec. index of slice, append item, delete all items, sort items in different ways)

Text operations (Functions, which are not in Go’s "strings" library, like "repeat" string n-times)

JSON operations (read JSON from file and transforming it in some way…, etc.)

Test your library functions, try to be as exhaustive as possible using test tables 

Think about useful abstractions modeled by interfaces in your library. An example could be that your library is reading in a  

file, so you imitate that behavior, or that you are doing a complex operation which should be mocked in the tests
