# awesome-stats

`awesome-stats` is a command line tool to modify an awesome list you maintain with various repository attributes. And with that it might help you choose the most popular, or most maintained, or quickly choose the non-abandoned tool or a library.

With default configuration the following

![Before](https://raw.githubusercontent.com/zerkms/awesome-stats/master/docs/before.png)

turns into

![After](https://raw.githubusercontent.com/zerkms/awesome-stats/master/docs/after.png)

# Usage

```
$ ./awesome-stats -h
Usage of ./awesome-stats:
  -in string
    	the input file
  -out string
    	the path to write output result
  -template string
    	stats template (default "(★{{.Stars}})")
```

# Download

You can download the recent build of the `awesome-stats` from https://bintray.com/zerkms/generic/awesome-stats/ or build it yourself (that would require go v1.8).

# License

GNU GENERAL PUBLIC LICENSE v3 (see `LICENSE.md`)
