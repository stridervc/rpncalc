rpncalc
=======
rpncalc is a simple command line RPN calculator.

RPN
---
RPN is short for reverse polish notation. 
An RPN calculator works with a stack. The numbers you enter
get pushed onto the stack. When you perform an operator, for
example '+', the last two numbers are popped from the stack, 
they're added together, and the result is popped onto the 
stack.

See https://en.wikipedia.org/wiki/Reverse_Polish_notation for
more details.

Usage
-----
Type a number and press <Enter> to push it onto the stack.
Other keys are
* '+' : Add
* '-' : Subtract
* '\*' : Multiply
* '/' : Divide
* '^' : Power of
* 's' : Swap the last two numbers on the stack
* 'd' : Drop the last number from the stack
* 'c' : Clear the stack completely
* 'n' : Negate the last number on the stack
