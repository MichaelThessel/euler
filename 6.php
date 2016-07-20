<?php

function find() {
    $sumOfSquares = 0;
    $sum = 0;
    for ($i = 1; $i <= 100; $i++) {
        $sumOfSquares += pow($i, 2);
        $sum += $i;
    }

    return pow($sum, 2) - $sumOfSquares;
}

var_dump(find()); die;
