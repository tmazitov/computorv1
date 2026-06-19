# computorv1

A command-line polynomial equation solver that handles equations up to degree 2, including complex roots.

## Dependencies

- Go 1.24+

## Usage

```bash
make
./computor "<equation>"
```

If no argument is provided, the program reads from stdin.

## Examples

```
./computor "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
0.905239
-0.475131
```

```
./computor "5 * X^0 + 4 * X^1 = 4 * X^0"
Reduced form: 1 * X^0 + 4 * X^1 = 0
Polynomial degree: 1
The solution is:
-0.25
```

```
./computor "1 * X^0 + 2 * X^1 + 5 * X^2 = 0"
Reduced form: 1 * X^0 + 2 * X^1 + 5 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly negative, the two complex solutions are:
-1/5 + 2i/5
-1/5 - 2i/5
```

```
./computor "6 * X^0 = 6 * X^0"
Reduced form: 0 * X^0 = 0
Any real number is a solution.
```
