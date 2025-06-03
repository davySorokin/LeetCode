/**
 * @param {number[]} status
 * @param {number[]} candies
 * @param {number[][]} keys
 * @param {number[][]} containedBoxes
 * @param {number[]} initialBoxes
 * @return {number}
 */
var maxCandies = function(status, candies, keys, containedBoxes, initialBoxes) {
    const n = status.length;
    const haveBox = new Array(n).fill(false);
    const haveKey = new Array(n).fill(false);
    const visited = new Array(n).fill(false);
    let queue = [...initialBoxes];
    let total = 0;

    initialBoxes.forEach(box => {
        haveBox[box] = true;
    });

    while (true) {
        let progress = false;
        const nextQueue = [];

        for (let box of queue) {
            if (visited[box] || (status[box] === 0 && !haveKey[box])) continue;

            visited[box] = true;
            total += candies[box];
            progress = true;

            for (let key of keys[box]) {
                if (!haveKey[key]) {
                    haveKey[key] = true;
                    if (haveBox[key] && !visited[key]) {
                        nextQueue.push(key);
                    }
                }
            }

            for (let b of containedBoxes[box]) {
                if (!haveBox[b]) {
                    haveBox[b] = true;
                    if (!visited[b]) {
                        nextQueue.push(b);
                    }
                }
            }
        }

        if (!progress) break;

        queue = nextQueue;
    }

    return total;
};
