<?php

$testData = [
    '(){}[]' => true,
    '([{}])' => true,
    '(}' => false,
    'not allowed' => false,
    '' => false,
    '}' => false,
    '({' => false,
    '[(])' => false,
    '[({})](]' => false,
    '(()' => false,
];

function validate(string $input): bool
{
    // validate if not empty
    if (empty($input)) {
        return error("String is empty!", $input);
    }

    // split string into symbols
    $splitted = str_split($input);

    // validate if symbol is brace
    $allowed = ['(', ')', '[', ']', '{', '}'];
    foreach ($splitted as $symbol) {
        if (!in_array($symbol, $allowed)) {
            return error("Symbol is not allowed!", $input);
        }

    }

    // validate first symbol is opened
    $opened = ['(', '{', '['];
    if (!in_array($splitted[0], $opened)) {
        return error("Not valid string! String must start with OPEN brace!", $input);
    }

    // validate last symbol is closed
    $closed = [')', '}', ']'];
    if (!in_array($splitted[count($splitted) - 1], $closed)) {
        return error("Not valid string! String must be closed with CLOSE brace!", $input);
    }

    $map = [
        '(' => ')',
        '{' => '}',
        '[' => ']'
    ];

    $counter = [];
    foreach ($map as $k => $v) {
        $counter[$k] = 0;
        $counter[$v] = 0;
    }

    // after OPEN brace should be also OPEN if not the same type
    // Ex.: (}
    foreach ($splitted as $k => $brace) {
        $counter[$brace]++;
        //if brace is NOT first
        if ($k > 0) {
            // if prev is OPEN
            if (in_array($splitted[$k - 1], $opened)) {
                // if brace is NOT opened
                if (!in_array($brace, $opened)) {
                    // throw err if not the same type
                    if ($map[$splitted[$k - 1]] !== $brace) {
                        return error("Can not put closed after opened with different types!", $input);
                    }
                }
            }
        }
    }
    foreach ($map as $k => $v) {
        if ($counter[$k] !== $counter[$v]) {
            return error("Not valid string!", $input);
        }
    }
    return true;
}

function error(string $message, string $input): bool
{
    echo $message . ". Input: " . $input . " \n";
    return false;
}

foreach ($testData as $input => $isValid) {
    $v = validate($input);
    if ($v) {
        echo "Valid! Input:" . $input . "\n";
    }
}