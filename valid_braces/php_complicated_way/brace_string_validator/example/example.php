<?php

require_once '../vendor/autoload.php';

$validator = \Brace\Factories\BracesStringValidatorFactory::create("(){}[]");
var_dump($validator->isBraceStringValid());

$validator = \Brace\Factories\BracesStringValidatorFactory::create("({)}[]");
var_dump($validator->isBraceStringValid());