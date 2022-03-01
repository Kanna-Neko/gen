## gen

A simple and subsidiary Command Line interface tool that can generate tests through generator file and solution file.

gen can help you ignore io problem in the generator file and solution file, you just need to output data to stdout in the generator file and output solution data like solve the problem.

if you don't need ans file, the solution file is unnecessary.

## example

### a typical sample to introduce basic funcion
```shell
tree
#.
#├── gen
#├── sampleGenerator.cpp
#└── sampleSolution.cpp
./gen sampleGenerator.cpp sampleSolution.cpp 3

#.
#├── gen
#├── generator
#├── sampleGenerator.cpp
#├── sampleSolution.cpp
#├── solution
#├── test1.in
#├── test1.out
#├── test2.in
#├── test2.out
#├── test3.in
#└── test3.out
```

### Flags

```code
1. -h, --help                  help for gen
2. -i, --inputSuffix string    add a suffix to all inputFile (default "in")
3. -o, --outputSuffix string   add a suffix to all outputFile (default "out")
4. -p, --prefix string         add a prefix to all fileName (default "test")
5. -s, --start int             set a starting sequence number before all files (default 1)
```