<?php

$number = 600851475143;

$i = 2;
while (true) {
    if ($number % $i == 0) {
        echo "Dividing by $i" . PHP_EOL;
        $number /= $i;
    }

    $i++;

    if ($i > $number) break;
}
