<?php

function generatePrimes($count) {
    $number = 9;
    $primes = [2, 3, 5, 7];
    do {
        $number += 2;

        if ($number % 2 == 0) continue;
        if ($number % 3 == 0) continue;
        if ($number % 5 == 0) continue;
        if ($number % 7 == 0) continue;

        $valid = true;
        foreach ($primes as $prime) {
            if ($number % $prime == 0) {
                $valid = false;
                break;
            }
        }

        if ($valid) $primes[] = $number;

    } while (count($primes) < $count);

    return array_pop($primes);
}

var_dump(generatePrimes(10001));
