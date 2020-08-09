<?php

namespace Brace\Interfaces;

use Brace\Entities\Brace;

/**
 * Interface IBraceFactory
 * @package Brace
 */
interface IBraceFactory
{
    public static function create(int $type, int $state): Brace;
}
