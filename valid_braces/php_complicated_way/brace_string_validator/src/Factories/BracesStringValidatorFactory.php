<?php

namespace Brace\Factories;

use Brace\Interfaces\IBraceStringValidator;
use Brace\Validators\BracesStringValidator;

class BracesStringValidatorFactory implements IBraceStringValidator
{
    /**
     * @param string $input
     * @return BracesStringValidator
     * @throws \Exception
     */
    public static function create(string $input): BracesStringValidator
    {
        return new BracesStringValidator($input);
    }
}