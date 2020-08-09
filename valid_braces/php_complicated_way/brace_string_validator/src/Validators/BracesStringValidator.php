<?php

namespace Brace\Validators;

use Brace\Constants\Constants;
use Brace\Entities\Brace;
use Brace\Enum\Enum;
use Brace\Factories\BraceFactory;
use Brace\Interfaces\IBraceFactory;

/**
 * Class Validator
 * @package Brace
 */
class BracesStringValidator
{
    /**
     * @var string
     */
    protected $input;

    /**
     * BraceValidator constructor.
     * @param string $input
     * @throws \Exception
     */
    public function __construct(string $input)
    {
        $this->input = $input;
    }

    /**
     * @return bool
     * @throws \Exception
     */
    public function isBraceStringValid(): bool
    {
        if (!$this->validateInput()) {
            return false;
        }

        $symbols = str_split($this->input);

        if (!$this->validateBracesSymbols($symbols)) {
            return false;
        }

        $braces = $this->getBraces($symbols);

        if (!$this->validateFirstOpenedBrace($braces)) {
            return false;
        }

        if (!$this->validateLastClosedBrace($braces)) {
            return false;
        }


        if (!$this->validateChain($braces)) {
            return false;
        }

        return true;
    }

    /**
     * @return bool
     */
    protected function validateInput(): bool
    {
        if (empty($this->input) || !is_string($this->input)) {
            return false;
        }
        return true;
    }

    /**
     * @param $braces
     * @return bool
     * @throws \Exception
     */
    protected function validateChain($braces): bool
    {
        foreach ($braces as $k => $v) {
            if (!$v instanceof Brace) {
                throw new \Exception("var is not instance of Brace/Entities/Brace");
            }
            if ($k == 0) {
                continue;
            }
            if (
                ($braces[$k - 1]->getState() === Enum::STATE_OPENED)
                && ($v->getState() === Enum::STATE_CLOSED)
                && ($braces[$k - 1]->getType() !== $v->getType())
            ) {
                return false;
            }
        }
        return true;
    }

    /**
     * @param array $braces
     * @return bool
     * @throws \Exception
     */
    protected function validateLastClosedBrace(array $braces): bool
    {
        $brace = $braces[count($braces) - 1];
        if (!$brace instanceof Brace) {
            throw new \Exception("var is not instance of Brace/Entities/Brace");
        }
        if ($brace->getState() === Enum::STATE_OPENED) {
            return false;
        }
        return true;
    }

    /**
     * @param array $braces
     * @return bool
     */
    protected function validateFirstOpenedBrace(array $braces): bool
    {
        if (!$braces[0] instanceof Brace) {
            throw new \Exception("var is not instance of Brace/Entities/Brace");
        }
        if ($braces[0]->getState() === Enum::STATE_CLOSED) {
            return false;
        }
        return true;
    }

    /**
     * @param array $symbols
     * @return bool
     */
    protected function validateBracesSymbols(array $symbols): bool
    {
        foreach ($symbols as $v) {
            if (!in_array($v, Constants::ALLOWED_SYMBOLS)) {
                return false;
            }
        }
        return true;
    }

    /**
     * @param array $symbols
     * @return array
     * @throws \Exception
     */
    protected function getBraces(array $symbols): array
    {
        $result = [];
        foreach ($symbols as $k => $v) {
            $result[] = BraceFactory::create(Constants::TYPES_MAP[$v], Constants::STATE_MAP[$v]);
        }
        return $result;
    }
}