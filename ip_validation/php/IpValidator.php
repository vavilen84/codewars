<?php

class IpValidator
{
    /** @var string */
    protected $ip;
    protected $pattern = '/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/';

    /**
     * @param string $ip
     */
    public function setIp(string $ip)
    {
        $this->ip = $ip;
    }

    public function validate(): bool
    {
        if (!is_string($this->ip)) {
            throw new \Exception('Ip must be string');
        }
        return (bool)preg_match($this->pattern, $this->ip);
    }
}