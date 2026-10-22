# Debugging notes

## 1. Accumulator variable is shadowed in `FromBase`

- **Input:** 101, from base 2.
- **What I saw:** After processing the first digit, the debugger showed `(n) = 0` and a separate `n = 1`. On the next loop iteration, the outer `n` was still `0`.
- **What I suspected:** The loop was creating a new local `n` instead of updating the accumulator declared before the loop.
- **How I confirmed it:** I stepped over the assignment and saw the inner `n` become `1` while the outer `(n)` stayed `0`. The function later returned the unchanged outer value.

## 2. `FromBase` accepts a digit equal to the base

- **Input:** `2`, from base 2.
- **What I saw:** The debugger showed `d = 2` while `base = 2`. After stepping over the validation condition, execution advanced to the accumulation line instead of returning `false`.
- **What I suspected:** The digit validation allowed a digit that is outside the valid range for the selected base.
- **How I confirmed it:** I followed execution past the condition with `d` equal to `base`. Since base 2 permits only digits 0 and 1, the input should have been rejected.

## 3. `ToBase` stops one digit too early

- **Input:** `10`, from base 2 to base 2 (decimal value 2).
- **What I saw:** In `ToBase`, the debugger showed `n = 2`, `base = 2`, and `result = "0"` after processing the first remainder. After dividing, `n` became `1`; execution then took the `break`, and the function returned `"0"`.
- **What I suspected:** The loop's stopping condition ended conversion before all digits had been added.
- **How I confirmed it:** I stepped through the division and condition, then saw the loop exit with `result = "0"`. The expected base-2 representation of decimal 2 is `"10"`.
