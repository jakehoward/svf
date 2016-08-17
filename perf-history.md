# Performance

## svf vs. cut ##

Test data: http://sdm.lbl.gov/fastbit/data/star2002-full.csv.gz

### v-alpha-0.1 ###

```
➜ time cut -d ',' -f 1,2,6 star2002-full.csv > /dev/null 
cut -d ',' -f 1,2,6 star2002-full.csv > /dev/null  49.83s user 1.55s system 97% cpu 52.502 total
```

```
➜ time svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null
svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null  122.44s user 11.05s system 107% cpu 2:04.58 total
```

