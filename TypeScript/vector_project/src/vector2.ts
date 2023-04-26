export default class Vector2 {
  constructor(public x: number, public y: number) {}

  add(vector: Vector2) {
    return new Vector2(vector.x + this.x, vector.y + this.y);
  }
}