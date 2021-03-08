class MyHashMap {
  array: number[]
  constructor() {
    this.array = []
  }

  put(key: number, value: number): void {
    this.array[key] = value
  }

  get(key: number): number {
    if (this.array[key] != undefined) return this.array[key]
    return -1
  }

  remove(key: number): void {
    this.array[key] = undefined
  }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * var obj = new MyHashMap()
 * obj.put(key,value)
 * var param_2 = obj.get(key)
 * obj.remove(key)
 */
