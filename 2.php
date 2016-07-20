<?php

$limit = 4000000;
$prev = 1;
$curr = 2;

$sum = 2;

while (true) {
    $tmp = $prev;
    $prev = $curr;
    $curr = $prev + $tmp;

    if ($curr > $limit) break;

    if ($curr % 2 == 0) $sum += $curr;
}

var_dump($sum); die;
