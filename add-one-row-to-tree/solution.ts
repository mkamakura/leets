/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function insert(v: number, node: TreeNode | null, d: number, n: number) {
  if (node === null) return
  if (d === n - 1) {
    const left = node.left
    node.left = new TreeNode(v, left, null)
    const right = node.right
    node.right = new TreeNode(v, null, right)
  } else {
    insert(v, node.left, d + 1, n)
    insert(v, node.right, d + 1, n)
  }
}

function addOneRow(root: TreeNode | null, v: number, d: number): TreeNode | null {
  if (d === 1) {
    const newTree = new TreeNode(v, root, null)
    return newTree
  }
  insert(v, root, 1, d)
  return root
}
