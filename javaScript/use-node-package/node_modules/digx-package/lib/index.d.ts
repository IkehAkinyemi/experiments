export default dig;
/**
 * A dig function that takes any object with a nested structure and a path,
 * and returns the value under that path or undefined when no value is found.
 *
 * @param {any}     source - A nested objects.
 * @param {string}  path - A path string, for example `my[1].test.field`
 * @param {boolean} [shouldThrow=false] - Optionally throw an exception when nothing found
 *
 */
declare function dig(source: any, path: string, shouldThrow?: boolean): any;
