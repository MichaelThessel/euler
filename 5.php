<?php

function find() {
    $number = 20;
    $divisors = range(11, 20);
    while (true) {
        $valid = true;
        foreach ($divisors as $divisor) {
            if ($number % $divisor != 0) {
                $valid = false;
                break;
            };
        }

        if ($valid) return $number;

        $number += 2;
    }
}

var_dump(find()); die;
