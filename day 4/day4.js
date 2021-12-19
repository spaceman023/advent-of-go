//js again. I will come back and do these in Go...eventually.
//parse input
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

//functions
function checkBoard(board) {
    //check row

    for (let i = 0; i < board.length; i++) {
        bingoCheck = true;
        for (let j = 0; j < board[i].length; j++) {
            let cell = board[i][j];
            if (cell !== true) {
                bingoCheck = false;
                break;
            }
        }
        if (bingoCheck === true) {
            console.log("bingo");
            return true;
        }
    }
    //check columns
    for (let i = 0; i < board[0].length; i++) {
        bingoCheck = true;
        for (let j = 0; j < board.length; j++) {
            let cell = board[j][i];
            if (cell !== true) {
                bingoCheck = false;
                break;
            }
        }
        if (bingoCheck === true) {
            return true;
            console.log("bingo");
        }
    }
    return false;
}
function markBoard(board, num) {
    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            let cell = board[i][j];
            if (cell === num) {
                board[i][j] = true;
            }
        }
    }
    return board;
}

function sumBoard(board) {
    sum = 0;
    for (let i = 0; i < board.length; i++) {
        for (let j = 0; j < board[i].length; j++) {
            let cell = board[i][j];
            if (cell !== true) {
                sum += cell;
            }
        }
    }
    return sum;
}

function solve(boards, moves) {
    sums = [];
    for (let i = 0; i < moves.length; i++) {
        move = moves[i];
        for (let j = 0; j < boards.length; j++) {
            let board = boards[j];
            //console.log(board);
            board = markBoard(board, move);
            //console.log(board);
            if (checkBoard(board)) {
                sums.push(sumBoard(board) * move);
                boards.splice(j, 1);
                j--;
            }
        }
    }
    return sums;
}
console.log(solve(boards, moves));
