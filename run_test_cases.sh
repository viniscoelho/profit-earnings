#!/bin/bash
for i in {0..8}
do
   ./profit-earnings < test_cases/test_case_$i.txt
done