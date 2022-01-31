# Wordle

## find

searches for words based on criteria

```
wordle find --help                                  
Usage:
  wordle find [flags]

Flags:
      --exclude string          letters to exclude
      --has string              letters in word
  -h, --help                    help for find
      --not-order stringArray   letters in the word but in wrong order (? is unknown)
      --order string            order of the word (? is unknown)

```

example

```
wordle find --exclude aroseunbz --has ilt --not-order '??til' --not-order '?lit?'
```

### exclude

letters that should be excluded from a word

```
--exclude asdf
```

### has

letters that should be included in a word

```
--has asdf
```

### no-order

letters that should be included but are in the wrong order. `?` denotes an unknown letter

```
-- no-order ??a?? --no-order ?a???
```

### order

letters that are included and in the correct order. `?` denotes an unknown letter

```
--order ??lk?
```