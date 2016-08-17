time svf -d ',' -f 1,2,6 star-2002-quarter.csv --cpuprofile=svf2.prof > /dev/null
go tool pprof /Users/jakehoward/Development/go/bin/svf svf2.prof
