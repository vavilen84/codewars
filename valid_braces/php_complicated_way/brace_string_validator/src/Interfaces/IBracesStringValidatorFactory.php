<?php

namespace Brace\Interfaces;

use Brace\Validators\BracesStringValidator;

/**
 * Interface IBraceStringValidator
 * @package Brace
 */
interface IBraceStringValidator
{
    public static function create(string $input): BracesStringValidator;
}
