#!/usr/bin/env awk

{
  print $0
  if (match($0, /^total:/)) {
    sub(/%/, "", $NF);
    printf("Test coverage is %s%% (Quality gate is %s%%)\n", $NF, target)
    if (strtonum($NF) < target) {
      printf("Test coverage does not meet expectations: %d%%, please add test cases!\n", target)
      exit 1;
    }
  }
}