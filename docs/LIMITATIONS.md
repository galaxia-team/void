# limitations
these are the current limitations of void. these may be fixed in the future.

limitations:
- comments cannot be on the same line as a piece of code
	* (bad) example 1:
        + `do_stuff() // example comment`
	* (good) example 2:
        + `// example comment`
- dictionary keys/values cannot contain a colon
	* (bad) example 1:
        + `dict = { "key": "value :" }`
	* (good) example 2:
        + `dict = { "key": "value" }`
- dictionary definitions cannot contain key/value pairs where the value is also a dictionary
    * note: you can (in the future) define a mutable dictionary with the key and an empty value, and then change the value to the sub-dictionary.
    * (bad) example 1:
        + `dict = { "testkey1": { "testkey2": "testvalue1" } }`
    * (good) example 2:
        + `dict = { "testkey1": "" }`
        + `dict["testkey1"] = { "testkey2": "testvalue1"}`
