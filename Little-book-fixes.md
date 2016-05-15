
+ ~~Fix the double-spacing issue for fmt.Println _(PR accepted)_~~
+ Functions section mentions named return values as if this has been covered in the text.
+ The use of 5 for index and 5 for value is confusing in slices
+ Pointers vs values for maps: reiterate the lessons learned there
+ Needs an actual example of how to compose interfaces.
+ Should probably have an actual example of the panic and recover functions
+ "Whatever you defer will be executed after the method returns" should be "Whatever you defer will be executed after the *enclosing* method returns" _(PR submitted)_
+ For the synchronisation output `i < 20` is better since it produces erratic output
+ Look into why `time.After` isn't classed as an "available channel" and randomly written to when `c` is also available.
