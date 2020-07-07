<?php

include_once "IpValidator.php";

$validator = new IpValidator();

$invalid = [
    '',
    'invalid input',
    '.1',
    '1.2.3',
    '1.2.3.4.',
    '1.2.3.4.5',
    '123.456.78.90',
    '123.045.067.089',
];

function validateInvalid(array $invalid, IpValidator $validator)
{
    foreach ($invalid as $ip) {
        $validator->setIp($ip);
        $isValid = $validator->validate();
        if ($isValid !== false) {
            throw new Exception("Failed!");
        }
    }
}

function validateValid(string $ip, IpValidator $validator)
{
    $validator->setIp('123.123.123.123');
    $isValid = $validator->validate();
    if ($isValid !== true) {
        throw new Exception("Failed!");
    }
}

validateInvalid($invalid, $validator);
validateValid('123.123.123.123', $validator);
echo 'OK!';