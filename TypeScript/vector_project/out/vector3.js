"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Vector3 = void 0;
class Vector3 {
    constructor(x, y, z) {
        this.x = x;
        this.y = y;
        this.z = z;
    }
    add(vector) {
        return new Vector3(vector.x + this.x, vector.y + this.y, vector.z + this.z);
    }
}
exports.Vector3 = Vector3;
