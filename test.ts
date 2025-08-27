// SOAL NO 1
function faktorial(n: number): number {
  if (n === 0) {
    return 1;
  }

  return n * faktorial(n - 1);
}

function hitungFaktorialDanBagi(n: number): number {
  return  Math.ceil(faktorial(n) / Math.pow(2, n));
}

console.log(hitungFaktorialDanBagi(0));
console.log(hitungFaktorialDanBagi(1));
console.log(hitungFaktorialDanBagi(2));
console.log(hitungFaktorialDanBagi(3));
console.log(hitungFaktorialDanBagi(4));
console.log(hitungFaktorialDanBagi(5));