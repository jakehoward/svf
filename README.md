# WIP: svf - separated values file #

Intended for use as a command line tool, similar to cut. `svf` is designed to read csv and tsv files so they can painlessly be used in bash pipelines.

## Examples ##

If you have a csv file that looks like:

```
census.csv
id,first\_name,last\_name
1,"Jake, the snake",Howard
2,Johnny,Dimblebert
...
```
tools like `cut` aren't well placed to handle the quoted second field, `svf` takes into account *sv escaping rules.

```
$ tail -n +2 census.csv | svf --csv -d ',' -f 2 | sort | uniq -c | sort -n > name\_popularity.txt
```
