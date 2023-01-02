function add(x: number): number {
  return x + 5;
}

function multiply(x: number): number {
  return x * 10;
}

function divide(x: number): number {
  return x / 2;
}

const func1 = (v1: string, v2: string) => `func1(${v1}, ${v2}, ...)`;
const func2 = (v: string) => `func2(${v})`;
const func3 = (v: string) => `func3(${v})`;


export const pipe = <T extends any[], R>(
  fn1: (...args: T) => R,
  ...fns: Array<(a: R) => R>
) => {
  const piped = fns.reduce((prevFn, nextFn) => (value: R) => nextFn(prevFn(value)), value => value);
  return (...args: T) => piped(fn1(...args));
};

const pipedFunction = pipe(func1, func2, func3);
console.log(pipedFunction("value1", "value2"));


// export const pipe = <T>(fn1: (a: T) => T, ...fns: Array<(a: T) => T>) =>
//   fns.reduce((prevFn, nextFn) => value => nextFn(prevFn(value)), fn1); 




// const compose = <T>(fn1: (a: T) => T, ...fns: Array<(a: T) => T>) =>
// fns.reduce((prevFn, nextFn) => value => prevFn(nextFn(value)), fn1); 

// const compose = <T extends any[], U>(
//   fn1: (...args: T) => U,
//   ...fns: Array<(x: U) => U>
// ) => {
//   const composed = fns.reduce((prevFn, nextFn) => (value: U) => prevFn(nextFn(value)), value => value);
//   return (...args: T) => composed(fn1(...args));
// }


// const composedFunction = compose(func1, func2, func3);
// console.log(composedFunction("value1", "value2"));