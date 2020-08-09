<?php

namespace Brace\Tests;

use PHPUnit\Framework\TestCase;

class BraceStringValidatorTest extends TestCase
{
    /**
     * @dataProvider additionProvider
     */
    public function testIsBraceStringValid(string $input, bool $expected)
    {
        $validator = \Brace\Factories\BracesStringValidatorFactory::create($input);
        $this->assertSame($expected, $validator->isBraceStringValid());
    }

    public function additionProvider()
    {
        return [
            ['(){}[]', true],
            ['([{}])', true],
            ['(}', false],
            ['not allowed', false],
            ['', false],
            ['}', false],
            ['({', false],
            ['[(])', false],
            ['[({})](]', false],
        ];
    }
}