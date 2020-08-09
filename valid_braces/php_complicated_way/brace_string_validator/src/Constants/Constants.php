<?php

namespace Brace\Constants;

use Brace\Enum\Enum;

/**
 * Class Constants
 * @package Brace
 */
class Constants
{
    const ALLOWED_SYMBOLS = ['(', ')', '[', ']', '{', '}'];

    const TYPES_MAP = [
        '(' => Enum::TYPE_ROUND,
        ')' => Enum::TYPE_ROUND,
        '[' => Enum::TYPE_SQUARE,
        ']' => Enum::TYPE_SQUARE,
        '{' => Enum::TYPE_CURLY,
        '}' => Enum::TYPE_CURLY,
    ];

    const STATE_MAP = [
        '(' => Enum::STATE_OPENED,
        ')' => Enum::STATE_CLOSED,
        '[' => Enum::STATE_OPENED,
        ']' => Enum::STATE_CLOSED,
        '{' => Enum::STATE_OPENED,
        '}' => Enum::STATE_CLOSED,
    ];
}