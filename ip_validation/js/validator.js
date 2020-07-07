class IpValidator {
    constructor() {
        this.pattern = '^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$';
    }

    isValidIp(ip) {
        let regExp = new RegExp(this.pattern);
        let result = ip.match(regExp);
        return result !== null;
    }
}

const validator = new IpValidator();

function testInvalid() {
    let result = true;
    let invalid = [
        '',
        'invalid input',
        '.1',
        '1.2.3',
        '1.2.3.4.',
        '1.2.3.4.5',
        '123.456.78.90',
        '123.045.067.089',
    ];

    for (let i = 0; i < invalid.length; i++) {
        let isValid = validator.isValidIp(invalid[i]);
        if (isValid) {
            console.log('Failed! Invalid case');
            result = false;
        }
    }
    return result;
}

function testValid() {
    let result = true;
    let isValid = validator.isValidIp('123.123.123.123');
    if (!isValid) {
        console.log('Failed! Valid case');
        result = false;
    }
    return result;
}

let resultInvalid = testInvalid();
let resultValid = testValid();
if (resultInvalid && resultValid) {
    console.log('OK!');
}

