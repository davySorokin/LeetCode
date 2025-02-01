/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */
var hammingDistance = function(x, y) {
    return (x ^ y).toString(2).split('').filter(bit => bit === '1').length;
};

// test cases
console.log(hammingDistance(1, 4)); // Output: 2
console.log(hammingDistance(3, 1)); // Output: 1
