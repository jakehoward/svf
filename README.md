# svf - separated value files #

`svf` is `cut` for separated value files. Designed to read csv and tsv files so they can be used in bash pipelines.

## Examples ##

If you have a csv file that looks like:

```
----------
census.csv
----------
id,first_name,last_name
1,"Jake, the snake",Howard
2,Johnny,Dimblebert
...
```
tools like `cut` aren't well placed to handle the quoted second field, `svf` takes into account *sv escaping rules.

```
$ tail -n +2 census.csv | svf -d , -f 2 | sort | uniq -c | sort -n > name_popularity.txt
```

## Workaround for missing column name feature (coming soon)
```
< data.csv | svf -d ',' -f $(< data.csv | python -c "import sys; x = sys.stdin.read(); print int(x.strip().split('COLUMN_NAME')[0].count(',')) + 1")
```
