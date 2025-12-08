const sum = require ('../src/index');
test ('sum 2+2 should be 4', () => {
    expect (sum(2,2)).toBe(4);
})