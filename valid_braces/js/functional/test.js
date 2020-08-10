let assert = require('assert');
let brace_validator = require("./brace_validator")

describe('Array', function () {
    describe('Validate isInputValid() method', function () {
        it('should return false: not allowed symbols ', function () {
            assert.equal(brace_validator.isInputValid("+"), false);
        });
        it('should return false: empty string', function () {
            assert.equal(brace_validator.isInputValid(""), false);
        });
        it('should return false: not allowed input type string', function () {
            assert.equal(brace_validator.isInputValid(["{}"]), false);
        });
    });
    describe('Validate isFirstSymbolOpen() method', function () {
        it('should return false', function () {
            assert.equal(brace_validator.isInputValid(")"), false);
        });
        it('should return true', function () {
            assert.equal(brace_validator.isInputValid("("), false);
        });
    });
    describe('Validate isFirstSymbolOpen() method', function () {
        it('should return false', function () {
            assert.equal(brace_validator.isInputValid(")"), false);
        });
        it('should return true', function () {
            assert.equal(brace_validator.isInputValid("("), false);
        });
    });
    describe('Validate isLastSymbolClose() method', function () {
        it('should return false', function () {
            assert.equal(brace_validator.isInputValid("(("), false);
        });
        it('should return true', function () {
            assert.equal(brace_validator.isInputValid("()"), false);
        });
    });
    describe('Validate isLastSymbolClose() method', function () {
        it('should return false', function () {
            assert.equal(brace_validator.isInputValid("(("), false);
        });
        it('should return true', function () {
            assert.equal(brace_validator.isInputValid("()"), false);
        });
    });
    describe('Validate isChainingValid() method', function () {
        it('should return false', function () {
            assert.equal(brace_validator.isInputValid("(()"), false);
        });
    });
});