import * as vectors from "./vectors";

const vec2a = new vectors.Vector2(1, 2);
const vec2b = new vectors.Vector2(2, 1);

console.log(vec2a.add(vec2b));

const vec3a = new vectors.Vector3(1, 2, 3);
const vec3b = new vectors.Vector3(3, 2, 1);

console.log(vec3a.add(vec3b));