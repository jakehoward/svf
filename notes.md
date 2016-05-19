# Notes

## *sv standards
- https://en.wikipedia.org/wiki/Delimiter-separated_values

### csv
- https://en.wikipedia.org/wiki/Comma-separated_values
- https://www.ietf.org/rfc/rfc4180.txt

### tsv
- https://en.wikipedia.org/wiki/Tab-separated_values
- http://www.iana.org/assignments/media-types/text/tab-separated-values
- http://www.cs.tut.fi/~jkorpela/TSV.html

## Test cases
- [ ] ,1,2,  3  ,"I'm a, nasty, csv file!",  "Jake ""the snake"" Howard"  ,"Dear sir,
This is a letter so has newlines...",2915,,,
- [ ] Multiple newines
- [ ] Missing lines
- [ ] Jagged csv
- [ ] Nonsense data (binary?)
- [ ] Different encodings of data/character types
- [ ] Security? 
- [ ] Extremely long lines
- [ ] Extremely large fields (possibly over multiple lines)
- [ ] Last line doesn't end in newline
- [ ] Windows line endings

## Design decisions
- Validation and errors?
- allow use of header name(s) as field selector(s)?
