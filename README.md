## gen

A simple and subsidiary Command Line interface tool that can generate tests through generator file and solution file.

gen can help you ignore io problem in the generator file and solution file, you just need to output data to stdout in the generator file and output solution data like solve the problem.

if you don't need ans file, the solution file is unnecessary.

## Usage

gen generateFileName solutionFileName [flags]

## example

### a typical sample to introduce basic funcion
```shell
tree
#.
#├── gen
#├── sampleGenerator.cpp
#└── sampleSolution.cpp
./gen sampleGenerator.cpp sampleSolution.cpp

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
3. -n, --num int               The number of input file and output file (default 10)
4. -o, --outputSuffix string   add a suffix to all outputFile (default "out")
5. -p, --prefix string         add a prefix to all fileName (default "test")
6. -s, --start int             set a starting sequence number before all files (default 1)
```

## develop todo List
- [x] Support cpp file
- [x] Functionalize the code
- [ ] Support python file (python3)
- [ ] Support c file
- [ ] Support golang file
- [ ] Support for generating compressed packages
- [x] generate outFiles through solution files and inputFiles, this function accomplish by command named gen solution