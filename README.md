## Intro
This is a response to the assignment provided by Liquid.

## Considerations:
- A monetary amount could be $4.00 or €2.17.
- Reports aggregate all of these amounts and present a total monetary sum.
- All present this data in Euros, often rounded to a whole number.
- The data will primarily come from accounting software or banking software, either through
  a file or an API integration.

## Task
Design a type (class / struct) that can handle these monetary amounts in a safe
way.

### Which operations would make sense for this/such type?
basics:
- addition
- subtraction
- multiplication
- division

complex:
- currency conversion
- modulus
- financial tax (for instance do certainNumber.Net() )

### How do we ensure that results from operations are accurate?
Integration and unit tests

### How do we handle rounding differences?


### How do we handle summing amounts of different currencies?
### Should this/such type be able to express negative amounts?
### How can I make sure that this/such type is easy to use correctly and hard to use incorrectly?



## Language - why Golang?


## Tests amd constraint verification
We would like to see that the code works and that you verify that it works correctly given the
constraints of this/such type.


## Future





https://github.com/shopspring/decimal
https://pkg.go.dev/math/big

