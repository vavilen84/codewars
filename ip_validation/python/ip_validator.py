import re
import unittest


def is_valid_ip(ip):
    pattern = '^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$'
    regexp = re.compile(pattern)
    return True if regexp.search(ip) else False


class TestIsValidIpMethods(unittest.TestCase):

    def test_is_valid_ip(self):
        invalid = [
            '',
            'invalid input',
            '.1',
            '1.2.3',
            '1.2.3.4.',
            '1.2.3.4.5',
            '123.456.78.90',
            '123.045.067.089',
        ]
        for ip in invalid:
            self.assertEqual(is_valid_ip(ip), False)

        self.assertEqual(is_valid_ip("123.123.123.123"), True)


if __name__ == '__main__':
    unittest.main()
