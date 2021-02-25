import chessboard from "./chessboard-1.0.0.js"

const onBoardChange = (o, n) => {
  const fen = document.querySelector('code')
  fen.innerHTML = Chessboard.objToFen(n)
}

let opts = {
  onChange: onBoardChange,
}
const b = Chessboard('b', opts)
b.start(true)

const newFen = (e) => {
  console.log(e)
  console.log(e.target.value)
  if (e.code === 'Enter') {
    b.position(e.target.value, true)
    e.target.value = ''
  }
}

const input = document.querySelector('input')
input.addEventListener('keyup', newFen)

