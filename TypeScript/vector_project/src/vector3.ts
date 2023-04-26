export default class Vector3 {
  constructor(public x: number, public y: number, public z: number) {}

  add(vector: Vector3) {
    return new Vector3(
      vector.x + this.x,
      vector.y + this.y,
      vector.z + this.z,
    )
  }
}