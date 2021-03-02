# Abstract Factory
Example of Abstract Factory design pattern in Go

This example shows how to generate different factories using the same interface, which then are used to process records in an abstract way.

Run the unit tests:
```
go test -v
```

Expected output:
```
=== RUN   TestVoiceSpainGermanyCdr
--- PASS: TestVoiceSpainGermanyCdr (0.00s)
=== RUN   TestVoiceSpainArgentinaCdr
--- PASS: TestVoiceSpainArgentinaCdr (0.00s)
=== RUN   TestSmsSpainGermanyCdr
--- PASS: TestSmsSpainGermanyCdr (0.00s)
=== RUN   TestSmsSpainArgentinaCdr
--- PASS: TestSmsSpainArgentinaCdr (0.00s)
PASS
ok  	go_abstract_factory	0.069s
```
