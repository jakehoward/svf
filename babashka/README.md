# svf

_Seperated value files_

## Deps
```
brew install borkdude/brew/babashka
```

## Usage

From directory containing svf.clj
```
bb -m svf --help
```

## Recipes
Get some cols from `;` delimited file:
```
< my-sc-csv.csv | bb -m svf -d ';' -f 2-
```

Get idxs of columns in big csv file:
```
< ~/tmp/big-csv.csv | bb -m svf --report | grep "My Discriminator" | cut -d ':' -f 2
```

