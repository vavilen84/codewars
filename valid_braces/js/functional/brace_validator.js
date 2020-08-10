const allowed = ["(", ")", "{", "}", "[", "]"];
const openBraces = ["(", "{", "["];
const closeBraces = [")", "}", "]"];

const curlyType = 1;
const roundType = 2;
const squareType = 3;

const types = [roundType, curlyType, squareType]

const typesMap = {
    "(": roundType,
    ")": roundType,
    "{": curlyType,
    "}": curlyType,
    "[": squareType,
    "]": squareType,
}

function isBraceStringValid(input) {
    if (!isInputValid(input)) {
        return false;
    }
    if (!isFirstSymbolOpen(input)) {
        return false;
    }
    if (!isLastSymbolClose(input)) {
        return false;
    }
    return isChainingValid(input);
}

function isChainingValid(input) {
    for (let i = 1; i < input.length; i++) {
        if (
            openBraces.includes(input[input.length - 1])
            && closeBraces.includes(input[i])
            && (typesMap[input[input.length - 1]] !== typesMap[input[i]])
        ) {
            return false;
        }
    }
    let counter = {
        open: new Map(),
        closed: new Map()
    }
    for (let i in input) {
        let count;
        if (openBraces.includes(input[i])) {
            count = counter.open.get(typesMap[i])
        } else {
            count = counter.closed.get(typesMap[i])
        }
        count++
        counter.open.set(typesMap[i], count)
    }
    for (let i in types) {
        if (counter.open.get(types[i]) !== counter.closed.get(types[i])) {
            return false;
        }
    }
    return true;
}

function isInputValid(input) {
    if (typeof input !== "string") {
        return false;
    }
    if (input.length < 1) {
        return false;
    }
    for (let i in input) {
        if (!allowed.includes(i)) {
            return false;
        }
    }
    return true;
}

function isFirstSymbolOpen(input) {
    return openBraces.includes(input[0]);

}

function isLastSymbolClose(input) {
    return closeBraces.includes(input[input.length - 1]);
}

module.exports = {
    isBraceStringValid: isBraceStringValid,
    isInputValid: isInputValid,
    isFirstSymbolOpen: isFirstSymbolOpen,
    isLastSymbolClose: isLastSymbolClose,
    isChainingValid: isChainingValid,
};
