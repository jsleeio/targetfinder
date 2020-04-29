# targetfinder

Some media organisations publish a "Target" word puzzle, where the objective is
to find as many words as possible using the set of letters provided, and where
every valid word must contain a particular letter.

[Here's a link to another solver](https://nineletterword.tompaton.com/ietrnmofd/).

The Unix (including macOS) system dictionary is used and assumed to be in
`/usr/share/dict/words`. Alternate dictionaries may be supplied.

## usage

```
$ ./targetfinder -help
Usage of ./targetfinder:
  -dictionary string
    	location of system dictionary file (default "/usr/share/dict/words")
  -key string
    	key letter. Must be exactly 1 character (default "t")
  -target string
    	target letters. Must be exactly 9 characters (default "dtgmtiaei")
```
