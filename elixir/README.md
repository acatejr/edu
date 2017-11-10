# Atoms Booleans and Strings

An atom is an immutable string.  
Simmilar to concept of symbols in Ruby.  
Written as `:hello`  
Used to reference  
A string that references itself.  It never changes.

## Booleans
* They are atoms.  

## Strings  
* UTF8 compliant  
* There is a difference between single and double quotes.  
* Single quotes = list of characters, a squence of bites.  It is binary.  

# Basics of Anonymous Functions and Tuples - Complex Types

## Anonymous functions - essentially a closure
`
fn(p) -> "hello #{p}" end
hello = fn(p) -> "hello #{p}" end
hello.('elixir')
"hello elixir"

hello = fn p -> "hello #{p}" end
`

## Tuple
An immutable indexed array.  
The elements are indexed.  
<code>
{1,2,3}  
elem({1,2,3}, 1)  
2  
tuple = {1,2,3}  
elem(tuple, 1)  
2
tuple_size(tuple)
3  
hello = fn -> {1,2,3} end
hello.()
{1,2,3}

</code>

## Pattern Matching  
Elixer uses '=' as a match operator.  
Joe Armstrong likes it to algebra x = y + 1  
a = 8 is binding.  8 is bound to a.  
8 = a is pattern matching.  

Pin operator "^".  
