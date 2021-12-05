//js again. I will come back and do these in Go...eventually.

let input = document.getElementsByTagName("pre")[0].innerText;
input = input.split("\n");
let moves = input[0].split(",").map((i) => Number(i));
let boards = [];
let currentboard = [];
for (let i = 2; i < input.length; i++) {
    if (input[i] == "") {
        continue;
    }
    if (currentboard.length == 5) {
        boards.push(currentboard);
        currentboard = [];
    }
    let build = input[i]
        .split(" ")
        .filter((i) => i !== "")
        .map((i) => Number(i));
    currentboard.push(build);
}
