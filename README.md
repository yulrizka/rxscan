# rxscan

rxscan provides functionality to scan text to variables using regular expression capture group.

This library is still experimental, use at your own risk. Contributions are always welcome and please
submit an issue if you find any problem.

## Examples

**Scanning a string**
```go
	input := "bright white bags contain 9 shiny gold bag."
	rx := regexp.MustCompile(`([\w ]+) bags contain (\d+) ([\w ]+) bag.`)
	var (
		bag1, bag2 string
		i          int
	)
	n, err := rxscan.Scan(rx, input, &bag1, &i, &bag2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("parsed %d arguments: %s -> (%d) %s", n, bag1, i, bag2)

	// Output: parsed 3 arguments: bright white -> (9) shiny gold
```

**Scanning repeated pattern**
```go
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	rx := regexp.MustCompile(`(\d+) ([\w ]+) bag`)

	sc := rxscan.NewScanner(rx, input)
	for sc.More() {
		var count int
		var color string
		_, err := sc.Scan(&count, &color)
		if err != nil {
			panic(err)
		}

		fmt.Printf("- (%d) %s\n", count, color)
	}
	if err := sc.Error(); err != nil {
		panic(err)
	}

	// Output:
	//- (1) bright white
	//- (2) muted yellow
	//- (3) bright white
	//- (4) muted yellow
	//- (1) shiny gold
	//- (2) shiny gold
	//- (9) faded blue
	//- (1) dark olive
	//- (2) vibrant plum
	//- (3) faded blue
	//- (4) dotted black
	//- (5) faded blue
	//- (6) dotted black
```
