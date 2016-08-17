# Performance

## svf vs. cut ##

Test data: http://sdm.lbl.gov/fastbit/data/star2002-full.csv.gz

```
➜ time cut -d ',' -f 1,2,6 star2002-full.csv > /dev/null 
cut -d ',' -f 1,2,6 star2002-full.csv > /dev/null  49.83s user 1.55s system 97% cpu 52.502 total
```

### v-alpha-0.1 ###

```
➜ time svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null
svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null  122.44s user 11.05s system 107% cpu 2:04.58 total
```

### v-alpha-0.2 ###

```
➜ time svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null       
svf -d ',' -f 1,2,6 star2002-full.csv > /dev/null  98.26s user 2.03s system 101% cpu 1:38.97 total
```
