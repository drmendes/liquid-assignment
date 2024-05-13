## Intro
This is a response to the assignment provided by Liquid.


# Challenges
_Need to improve the stability and scalability relatively fast._

Regarding scalability and stability, I believe the most crucial piece is having a fast and smooth integration with third party tools clients 
rely upon.

_Improve our data security_

The most effective strategy to enhance data security includes:
- Implementing strict access controls, role-based permissions, and adherence to the principle of least privilege.
- Employing encryption to allow clients exclusive access to their data through keys.
- Encrypting data both in transit and at rest.
- Conducting regular security audits to preempt potential vulnerabilities.

# Technical assignment

## Task
Design a type (class / struct) that can handle these monetary amounts in a safe way.

## Considerations:
- A monetary amount could be $4.00 or €2.17.
- Reports aggregate all of these amounts and present a total monetary sum.
- All present this data in Euros, often rounded to a whole number.
- The data will primarily come from accounting software or banking software, either through
a file or an API integration.


## Questions
- Which operations would make sense for this/such type?

Basics such as addition, subtraction, multiplication, division, rounding, absolute value, etc.
In case we choose to have our own representation of money, we can implement more complex ones such as currency 
conversion and tax calculation.

- How do we ensure that results from operations are accurate?

In this kind of sensitive data, its essential to cover all possible scenarios with unit testing. 

- How do we handle rounding differences?

A rounding strategy must be agreed upon to be followed. Using _big.Int_ allows us to all the precision we need.
Problems arise with irrational numbers such as 1/3. Here standards need to be defined.

- How do we handle summing amounts of different currencies?

Given we want to represent the sum in currency A, we should just apply the currency exchange rate to the second 
amount and make the normal sum.

- Should this/such type be able to express negative amounts?

Yes, they are quite common and their lack could become a problem with third party integrations.

- How can I make sure that this/such type is easy to use correctly and hard to use
  incorrectly?

By providing clear methods for common operations and leveraging Golang’s interfaces and type safety.


## Implementation
I ended up making something heavily related to decimal.Decimal from [Shopspring package](https://github.com/shopspring/decimal). 
It makes use of [big.Int](https://pkg.go.dev/math/big) and basically is a rational number where the denominator is always a power of 10. 
This may, however, be considered overkill over other approaches.
In a production environment I would use the package itself or a fork of it.



## Language - why Golang?
A bit generic but:
 
Pros:
 
- Performant (which is a must in the financial realm)
- Built-in concurrency support with goroutines and channels, ideal for processing large datasets concurrently.
- Statically typed, which helps catch errors at compile time rather than runtime
- Compiles to a single binary, making deployment straightforward and reducing dependency issues.

Cons:

- Error handling in Golang can be verbose due to the lack of exceptions or advanced Result types.
- Library Ecosystem is not has big and established as, for instance, Javascript or Python.

