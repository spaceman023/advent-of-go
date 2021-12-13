
// Importing the fs module
let fs = require("fs")
// Intitializing the readFileLines with the file
const readFileLines = filename =>
    fs.readFileSync(filename)
        .toString('UTF8')
        .split('\n');

// Calling the readFiles function with file name
let arr = readFileLines('../inputs/input12sample.txt');
// Printing the response array
console.log(arr);
let v;
let adjList;
// A directed graph using
// adjacency list representation
function Graph(vertices) {
    // initialise vertex count
    v = vertices;
    // initialise adjacency list
    initAdjList();
}
// utility method to initialise
// adjacency list
function initAdjList() {
    adjList = new Array(v);
    for (let i = 0; i < v; i++) {
        adjList[i] = [];
    }
}
// add edge from u to v
function addEdge(u, v) {
    // Add v to u's list.
    adjList[u].push(v);
}
// Prints all paths from
// 's' to 'd'
function printAllPaths(s, d) {
    let isVisited = new Array(v);
    for (let i = 0; i < v; i++)
        isVisited[i] = false;
    let pathList = [];

    // add source to path[]
    pathList.push(s);

    // Call recursive utility
    printAllPathsUtil(s, d, isVisited, pathList);
}
// A recursive function to print
// all paths from 'u' to 'd'.
// isVisited[] keeps track of
// vertices in current path.
// localPathList<> stores actual
// vertices in the current path
let count = 0
function printAllPathsUtil(u, d, isVisited, localPathList) {
    if (u == (d)) {
        count++
        console.log((localPathList + "<br>"));
        // if match found then no need to
        // traverse more till depth
        return;
    }
    // Mark the current node
    if (nodes[u] === nodes[u].toLowerCase()) {
        isVisited[u] = true;
    }
    // Recur for all the vertices
    // adjacent to current vertex
    for (let i = 0; i < adjList[u].length; i++) {
        if (!isVisited[adjList[u][i]]) {
            // store current node
            // in path[]
            localPathList.push(adjList[u][i]);
            printAllPathsUtil(adjList[u][i], d,
                isVisited, localPathList);
            // remove current node
            // in path[]
            localPathList.splice(localPathList.indexOf
                (adjList[u][i]), 1);
        }
    }
    // Mark the current node
    isVisited[u] = false;
}
// Driver program
// Create a sample graph
Graph(100);
let nodes = new Set()
for (let vert of arr) {
    vert = vert.split('-')
    nodes.add(vert[0])
    nodes.add(vert[1])
}
nodes = Array.from(nodes)
for (let vert of arr) {
    vert = vert.split('-')
    id1 = nodes.indexOf(vert[0])
    id2 = nodes.indexOf(vert[1])
    addEdge(id1, id2)
    addEdge(id2, id1)
}

// arbitrary source
let s = nodes.indexOf('start');
// arbitrary destination
let d = nodes.indexOf('end');
console.log(
    "Following are all different paths from "
    + s + " to " + d + "<Br>");
printAllPaths(s, d);
console.log(count)
