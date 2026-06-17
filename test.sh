#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

BINARY="./computor"
PASS=0
FAIL=0

run_test() {
	local description="$1"
	local input="$2"
	local expected="$3"

	actual=$("$BINARY" "$input" 2>&1)
	if [ "$actual" = "$expected" ]; then
		printf "${GREEN}PASS${NC} %s\n" "$description"
		PASS=$((PASS + 1))
	else
		printf "${RED}FAIL${NC} %s\n" "$description"
		diff <(echo "$expected") <(echo "$actual") | sed 's/^/       /'
		FAIL=$((FAIL + 1))
	fi
}

make > /dev/null 2>&1 || { echo "Build failed"; exit 1; }

# --- subject examples ---

run_test "degree 2, d > 0, two real roots" \
	"5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0" \
	"Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
0.905239
-0.475131"

run_test "degree 1, one real root" \
	"5 * X^0 + 4 * X^1 = 4 * X^0" \
	"Reduced form: 1 * X^0 + 4 * X^1 = 0
Polynomial degree: 1
The solution is:
-0.25"

run_test "degree 3, can't solve" \
	"8 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 3 * X^0" \
	"Reduced form: 5 * X^0 - 6 * X^1 + 0 * X^2 - 5.6 * X^3 = 0
Polynomial degree: 3
The polynomial degree is strictly greater than 2, I can't solve."

run_test "degree 0, any real number is a solution" \
	"6 * X^0 = 6 * X^0" \
	"Reduced form: 0 * X^0 = 0
Any real number is a solution."

run_test "degree 0, no solution" \
	"10 * X^0 = 15 * X^0" \
	"Reduced form: -5 * X^0 = 0
No solution."

run_test "degree 2, d < 0, two complex roots" \
	"1 * X^0 + 2 * X^1 + 5 * X^2 = 0" \
	"Reduced form: 1 * X^0 + 2 * X^1 + 5 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly negative, the two complex solutions are:
-1/5 + 2i/5
-1/5 - 2i/5"

# --- extra ---

run_test "degree 2, d = 0, one root" \
	"1 * X^0 - 2 * X^1 + 1 * X^2 = 0" \
	"Reduced form: 1 * X^0 - 2 * X^1 + 1 * X^2 = 0
Polynomial degree: 2
Discriminant is zero, the solution is:
1"

run_test "degree 1, negative coefficient" \
	"4 * X^0 - 2 * X^1 = 0" \
	"Reduced form: 4 * X^0 - 2 * X^1 = 0
Polynomial degree: 1
The solution is:
2"

echo ""
echo "Results: $PASS passed, $FAIL failed"
[ $FAIL -eq 0 ] && exit 0 || exit 1
