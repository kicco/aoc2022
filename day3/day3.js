const fs = require("fs")
const input = fs.readFileSync("input.txt", "utf-8")
let answer = 0

const priority = [
  undefined,
  "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"
]

input.split(/\r?\n/).forEach(line => {
  let half = line.length / 2
  let first = line.substring(0, half).split("")
  let second = line.substring(half, line.length).split("")
  for (i = 0; i < first.length; i++) {
    if (second.includes(first[i])) {
      answer += priority.indexOf(first[i])
      break;
    }
  }
})
console.log("Answer:", answer)
