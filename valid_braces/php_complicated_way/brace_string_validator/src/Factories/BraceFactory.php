<?php

namespace Brace\Factories;

use Brace\Interfaces\IBraceFactory;
use Brace\Entities\Brace;

class BraceFactory implements IBraceFactory
{
    /**
     * @param int $type
     * @param int $state
     * @return Brace
     * @throws \Exception
     */
    public static function create(int $type, int $state): Brace
    {
        self::validateInputData($type, $state);
        return self::getNewBrace($type, $state);
    }

    /**
     * @param $type
     * @param $state
     * @return Brace
     */
    protected static function getNewBrace($type, $state): Brace
    {
        $brace = new Brace();
        $brace->setState($state);
        $brace->setType($type);

        return $brace;
    }

    /**
     * @param int $type
     * @param int $state
     * @throws \Exception
     */
    protected static function validateInputData(int $type, int $state)
    {
        if (empty($type)) {
            throw new \Exception("type is required!");
        }
        if (empty($state)) {
            throw new \Exception("state is required!");
        }
    }
}