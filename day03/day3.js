//I did this in JS because I'm lazy

let input = document.getElementsByTagName("pre")[0].innerText;
input = input.split("\n");
gamma = "";
epsilon = "";
for (let j = 0; j < input[0].length; j++) {
    let ones = 0;
    let zeroes = 0;
    for (let i = 0; i < input.length; i++) {
        input[i][j] == 1 ? ones++ : zeroes++;
    }
    ones > zeroes ? (gamma += "1") : (gamma += "0");
    ones > zeroes ? (epsilon += "0") : (epsilon += "1");
}
gnum = parseInt(gamma, 2);
fnum = parseInt(epsilon, 2);
//part two
gcopy = [...input];
ecopy = [...input];
for (let i = 0; i < 12; i++) {
    if (gcopy.length === 1) {
        break;
    }
    let ones = 0;
    let zeroes = 0;
    for (let j = 0; j < gcopy.length; j++) {
        gcopy[j][i] === "1" ? ones++ : zeroes++;
    }
    if (ones >= zeroes) {
        gcopy = gcopy.filter((el) => el[i] === "1");
    }
    if (zeroes > ones) {
        gcopy = gcopy.filter((el) => el[i] === "0");
    }
}
for (let i = 0; i < 12; i++) {
    if (ecopy.length === 1) {
        break;
    }
    let ones = 0;
    let zeroes = 0;
    for (let j = 0; j < ecopy.length; j++) {
        ecopy[j][i] === "1" ? ones++ : zeroes++;
    }
    if (ones >= zeroes) {
        ecopy = ecopy.filter((el) => el[i] === "0");
    }
    if (zeroes > ones) {
        ecopy = ecopy.filter((el) => el[i] == "1");
    }
}
console.log(gcopy, ecopy);
ox = parseInt(gcopy[0], 2);
co = parseInt(ecopy[0], 2);
//console.log(co, ox)
console.log(ox * co);
