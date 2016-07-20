<?php

function isPalindrome($number) {
    $digits = getDigits($number);
    $numDigits = count($digits);
    for ($i = 0; $i <= ($numDigits / 2) - 1; $i++) {
        if ($digits[$i] != $digits[$numDigits - $i - 1]) return false;
    }

    return true;
}

function getDigits($number) {
    $digits = [];

    $mag = getMagnitude($number);
    for ($i = $mag; $i >= 0; $i--) {
        $div = pow(10, $i);
        $digits[$i] = bcdiv($number, $div, 0);
        $number -= $digits[$i] * $div;
    }

    return $digits;
}

function getMagnitude($number) {
    if ($number < 0) $number *= -1;

    for ($mag = 0; true; $mag++) {
        if (pow(10, $mag) > $number) {
            return $mag - 1;
        }
    }
}

function findPalindrome() {
    $palindrome = array(
        'product' => 0,
        'a' => 0,
        'b' => 0
    );

    for ($i = 999; $i >= 100; $i--) {
        for ($j = 999; $j >= 100; $j--) {
            $prod = $i * $j;
            if ($prod > $palindrome['product'] && isPalindrome($prod)) {
                $palindrome['product'] = $prod;
                $palindrome['a'] = $i;
                $palindrome['b'] = $j;
            }
        }
    }

    return $palindrome;
}

var_dump(findPalindrome()); die;
